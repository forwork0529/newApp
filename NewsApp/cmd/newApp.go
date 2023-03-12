package main

import (
	"newsApp/packages/filesIO"
	"newsApp/packages/logger"
	"newsApp/packages/rssReader"
	"time"
)

func main(){

	logs := logger.New()
	r := rssReader.New(filesIO.AConf, logs)
	r.Start()



}





