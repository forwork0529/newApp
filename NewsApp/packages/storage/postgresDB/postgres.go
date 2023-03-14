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
	pulse int
	db *pgxpool.Pool
	l *log.Logger
}


func New(ch  <- chan proxyStructs.Article, config proxyStructs.AppConfig, logs *log.Logger)(*postgresDB, error){

	conPool, err :=  pgxpool.Connect(context.Background(), config.ConnString)
	if err != nil{
		return nil, err
	}
	logs.Println("connected to postgresql..")
	return &postgresDB{
		ch : ch,
		pulse : config.ReqPer,
		db : conPool,
		l : logs,
	}, nil
}

func (p *postgresDB) Push(){
	var c int
	for{
		select{
		case <- p.ch: c += 1
			continue
		case <- time.After(time.Second * 3):
			p.l.Printf("already read from chan %v articles", c)
			return
		}
	}
}

func (p *postgresDB) Get(number int)([]proxyStructs.Article, error){
	return []proxyStructs.Article{}, nil
}


func (p *postgresDB)Start(){
	t := time.NewTicker(time.Duration(p.pulse))
	go func(){
		for{
			<- t.C
			p.Push()
		}
	}()
}
