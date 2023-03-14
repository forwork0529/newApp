package MemoryDB

import (
	"errors"
	"log"
	"newsApp/packages/proxyStructs"
	"time"
)

type memoryDB struct{
	ch <- chan  proxyStructs.Article
	pulse int
	l *log.Logger
}

func New(ch <- chan proxyStructs.Article, config proxyStructs.AppConfig, logs *log.Logger)*memoryDB{
	return &memoryDB{
		ch : ch,
		pulse : config.ReqPer,
		l : logs,
	}
}

func (m *memoryDB) Push(){
	var c int
	for{
		select{
		case <- m.ch: c += 1
		continue
		case <- time.After(time.Second * 3):
			m.l.Printf("already read from chan %v articles", c)
			return
		}
	}
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


var attendantArt  = proxyStructs.Article{
	Title: "Attendant article",
	Description: "Some description it the body of article",
	PubDate: "Fri, 10 Mar 2023 00:00:00 +0000",
	Link : "https://golangweekly.com/issues/451",
}

func (m *memoryDB)Start(){
	t := time.NewTicker(time.Duration(m.pulse))
	go func(){
		for{
			<- t.C
			m.Push()
		}
	}()
}

