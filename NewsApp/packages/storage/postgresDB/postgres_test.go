package postgresDB

import (
	"context"
	"log"
	"newsApp/packages/proxyStructs"
	"testing"
)

func TestNew(t *testing.T) {

	chFromNet := make(chan proxyStructs.Article, 30)
	config := proxyStructs.AppConfig{
		Sources : []string {""},
		ReqPer  : 3,
		BDType	: "",
		ConnString : "postgres://postgres:postgres@127.0.0.1:8081/articles",
	}
	logs := log.Default()

	db , err := New(chFromNet, config, logs)
	if err != nil{
		t.Fatal(err)
	}
	logs.Println("DB created")
	fullChan(chFromNet)
	logs.Println("the channel is filled by the program")
	t.Run("func_Push", func(t *testing.T) {
		c := 0
		rows, err := db.db.Query(context.Background(),`SELECT * FROM articles;`)
		if err != nil{
			t.Fatal(err)
		}
		defer rows.Close()

		for rows.Next(){
			c += 1
		}
		db.Push(proxyStructs.Article{
			ID : 0,
			Title: "Hi",
			Description : "World",
			PubDate: 123456,
			Link: "www",
		})
		if db.err != nil{
			t.Fatal(err)
		}
		db.Push(testAerticle)
		if db.err != nil{
			t.Fatal(err)
		}
		d := 0
		rows, err = db.db.Query(context.Background(),`SELECT * FROM articles;`)
		if err != nil{
			t.Fatal(err)
		}
		defer rows.Close()

		for rows.Next(){
			d += 1
		}
		if (d - 2) != c{
			t.Fatal("no new data after push (test can fail when the article with test title exists, sometimes need retry)")
		}

	})
	t.Run("func_Get", func(t *testing.T) {
		rows, err := db.db.Query (context.Background(), `SELECT * FROM ARTICLES;`)
		if err != nil{
			t.Fatal(err)
		}
		defer rows.Close()
		c := 0
		for rows.Next(){
			c += 1
		}
		db.l.Println(c)
		if c == 0 {
			t.Fatal("Not enough data for testing")
		}
		var list []proxyStructs.Article
		n := 3
		list, err = db.Get(n)
		if n >= c{
			if len(list) < c{
				t.Fatal("Not all data is extracted")
			}
		} else { if len(list) < n{
			t.Fatal("Not all data is extracted")
		}
		if list[0].PubDate < 2 || len(list[0].Title) < 2 || len(list[0].Description) < 2 || len(list[0].Link) < 2{
			t.Fatal("Incorrect data was extracted")
		}

		}
	})
	t.Run("func_Flush", func(t *testing.T) {
		rows, err := db.db.Query(context.Background(), `SELECT * FROM articles;`)
		if err != nil{
			log.Fatal(err)
		}
		defer rows.Close()
		c := 0
		for rows.Next(){
			c += 1
		}
		if c == 0{
			t.Fatal("Not enough data for testing")
		}
		db.Flush()

		c = 0
		rows, err = db.db.Query(context.Background(), `SELECT * FROM articles;`)
		if err != nil{
			log.Fatal(err)
		}
		defer rows.Close()
		for rows.Next(){
			c += 1
		}
		if c > 0{
			t.Fatal("flush func doesnt work")
		}
	})

}

func fullChan(ch chan proxyStructs.Article){
	for i := 1; i <= 10; i++{
		ch <- proxyStructs.Article{}
	}
}
var testAerticle = proxyStructs.Article{
	Title: "[Перевод] 5 паттернов параллельного программирования в GO, которые сделают ваш следующий проект лучше",
	Description: "<p>Параллельное программирование — одна из самых интересных фич, которые может предложить вам Golang. Идея, лежащая в основе параллелизма, заключается в одновременной работе над несколькими разными процессами, что помогает избежать застревания в задачах, выполнение которых занимает много времени. </p><p></p> <a href=\"https://habr.com/ru/post/722880/?utm_campaign=722880&amp;utm_source=habrahabr&amp;utm_medium=rss#habracut\">Читать далее</a>\n",
	PubDate: 1678976102,
	Link: "https://habr.com/ru/post/722880/?utm_campaign=722880&utm_source=habrahabr&utm_medium=rss\n",
}

/*
Тестировочная статья для загрузки непосредственно в psql..
INSERT INTO articles (title, content, published_at, link) VALUES ('[Перевод] 5 паттернов параллельного программирования в GO, которые сделают ваш следующий проект лучше',
'<p>Параллельное программирование — одна из самых интересных фич, которые может предложить вам Golang. Идея, лежащая в основе параллелизма, заключается в одновременной работе над несколькими разными процессами, что помогает избежать застревания в задачах, выполнение которых занимает много времени. </p><p></p> <a href=\"https://habr.com/ru/post/722880/?utm_campaign=722880&amp;utm_source=habrahabr&amp;utm_medium=rss#habracut\">Читать далее</a>\n',
1678976102, 'https://habr.com/ru/post/722880/?utm_campaign=722880&utm_source=habrahabr&utm_medium=rss');
*/

