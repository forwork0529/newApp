package postgresDB

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"newsApp/packages/proxyStructs"
	"time"
)

type postgresDB struct{
	ch <- chan proxyStructs.Article
	db *pgxpool.Pool
	l *log.Logger
	err error
	pulse int
}

// Создание структуры предназначенной для: регулярного получения , сохранения
// удаления и выдачи публикаций из БД. Выдача осуществляется методом Get,
// регулярность опредеяется полем струектуры конфигурации (требуется запуск).

func New(ch  <- chan proxyStructs.Article, config proxyStructs.AppConfig, logs *log.Logger)(*postgresDB, error){


	conPool, err :=  pgxpool.Connect(context.Background(), config.ConnString)
	if err != nil{
		return nil, err
	}
	err = conPool.Ping(context.Background())
	if err != nil{
		return nil, err
	}
	logs.Println("connected to postgresql..")
	return &postgresDB{
		ch : ch,
		db : conPool,
		l : logs,
		pulse : config.ReqPer,
	}, nil
}

// получение данных из структуры взаимодействующей с БД.
func (p *postgresDB) Get(number int)([]proxyStructs.Article, error){
	var res = make([]proxyStructs.Article, 0)
	rows, err := p.db.Query(context.Background(),`SELECT title, content, published_at, link FROM articles ORDER BY published_at DESC LIMIT $1`, number)
	if err != nil {
		return []proxyStructs.Article{}, err
	}
	defer rows.Close()
	for rows.Next(){
		var article proxyStructs.Article
		err = rows.Scan(&article.Title, &article.Description, &article.PubDate, &article.Link)
		if err != nil{
			return []proxyStructs.Article{}, err
		}
		res = append(res, article)
	}
	if rows.Err() != nil{
		return []proxyStructs.Article{}, err
	}
	return res, nil
}


// Добавление публикации в БД, если публикация с таким названием уже
// имеется в БД никаких действий не происходит
func (p * postgresDB) Push(art proxyStructs.Article){

	_, err := p.db.Exec(context.Background(), `INSERT INTO articles (title, content, published_at, link)
	VALUES ($1, $2, $3, $4) returning id;`, art.Title, art.Description, art.PubDate, art.Link)
/*	var id int
	err := p.db.QueryRow(context.Background(),`INSERT INTO articles (title, content, published_at, link)
	VALUES ($1, $2, $3, $4) returning id;`, art.Title, art.Description, art.PubDate, art.Link).Scan(&id)*/
	if err != nil{
		if err.Error() == "ERROR: duplicate key value violates unique constraint \"articles_title_key\" (SQLSTATE 23505)"{

		} else{
			p.updtErr(err)
		}

	}
}

// Полная очистка таблицы базы данных
func (p *postgresDB) Flush(){
	_, err := p.db.Exec(context.Background(), `truncate articles;`)
	if err != nil{
		p.l.Printf("error via truncating table: %v\n", err.Error())
	}
}

// Start (запуск) запускает циклическое чтение, логирование ошибок и удаление данных из
// структуры работающей с БД. Регулярность определяется полем конфигурационной структуры при создании
func (p *postgresDB)Start(){
	go func(){
		for{
			article := <- p.ch
			p.Push(article)

		}
	}()
	go func(){
		t := time.NewTicker(time.Duration(p.pulse) * time.Minute)
		for{
			<- t.C
			if p.err != nil{
				p.l.Printf("one of the storage errors: %v", p.err.Error())
			}
		}
	}()
	go func(){

		for{
			<- time.After(time.Duration(p.pulse) * time.Minute * 10)
			p.Flush()
		}
	}()
}


func (p * postgresDB) updtErr(err error){
	if p.err == nil{
		p.err = err
	}
}