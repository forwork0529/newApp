package main

import (
	"log"
	"net/http"
	"newsApp/packages/api"
	"newsApp/packages/filesIO"
	"newsApp/packages/logger"
	"newsApp/packages/rssReader"
	"newsApp/packages/storage"
)


func main(){

	logs := logger.New()
	r := rssReader.New(filesIO.AConf, logs)
	r.Start()


	db := storage.New(r.Chan(), filesIO.AConf, logs)

	db.Start()
	myAPI := api.New(db, logs)
	log.Fatal(http.ListenAndServe(":9090", myAPI.Router()))
}

