package MemoryDB

import (
	"errors"
	"log"
	"newsApp/packages/proxyStructs"
)

type memoryDB struct{
	ch <- chan  proxyStructs.Article
	l *log.Logger
}

func New(ch <- chan proxyStructs.Article, config proxyStructs.AppConfig, logs *log.Logger)*memoryDB{
	logs.Println("using memoryDB...")
	return &memoryDB{
		ch : ch,
		l : logs,
	}
}

func (m *memoryDB) Push(art proxyStructs.Article){

}

func (m *memoryDB) Get(number int)([]proxyStructs.Article, error){
	if number < 1{
		return []proxyStructs.Article{}, errors.New("number of articles must be > 0")
	}
	res := make([]proxyStructs.Article, 0)
	for i := 1; i <= number; i++{
		res = append(res, attendantArt)
	}
	return res, nil
}

func (m *memoryDB) Flush(){

}



func (m *memoryDB)Start(){
	go func(){
		for{
			article := <- m.ch
			m.Push(article)
		}
	}()
}


var attendantArt  = proxyStructs.Article{
	Title: "Attendant article",
	Description: "Some description it the body of article",
	PubDate: 123456,
	Link : "https://golangweekly.com/issues/451",
}