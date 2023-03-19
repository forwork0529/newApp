package MemoryDB

import (
	"log"
	"newsApp/packages/proxyStructs"
	"testing"
	"time"
)

func Test_postgresDB(t *testing.T) {

	chFromNet := make(chan proxyStructs.Article, 30)
	config := proxyStructs.AppConfig{
		Sources : []string {""},
		ReqPer  : 3,
		BDType	: "",
		ConnString : "",
	}
	logs := log.Default()

	db := New(chFromNet, config, logs)
	logs.Println("DB created")
	fullChan(chFromNet)
	logs.Println("the channel is filled by the program")
	db.Start()
	logs.Println("DB process started")
	time.Sleep(time.Second * 5)
	select{
	case <- chFromNet: t.Fatal("DB doesnt read from chan ..")
	case <- time.After(time.Second * 2): break
	}
	logs.Println("input channel is empty")

	res, err := db.Get(10)
	if len(res) != 10 || err != nil{
		t.Fatal("wrong data from get method")
	}

}

func fullChan(ch chan proxyStructs.Article){
	for i := 1; i <= 10; i++{
		ch <- proxyStructs.Article{}
	}
}