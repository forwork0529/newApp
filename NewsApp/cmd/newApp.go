package main

import (
	"newsApp/packages/filesIO"
	"newsApp/packages/logger"
	"newsApp/packages/rssReader"
	"newsApp/packages/storage"
	"time"
)

func main(){


	logs := logger.New()
	r := rssReader.New(filesIO.AConf, logs)
	r.Start()
	db := storage.New(r.Chan(), filesIO.AConf, logs)
	db.Start()

 	time.Sleep(time.Second * 10)

}





