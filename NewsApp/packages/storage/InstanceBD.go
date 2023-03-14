package storage

import (
	"log"
	"newsApp/packages/proxyStructs"
	"newsApp/packages/storage/MemoryDB"
	"newsApp/packages/storage/postgresDB"
)

type db interface{
	Push()
	Get(number int)([]proxyStructs.Article, error)
	Start()
}

func New(ch <- chan proxyStructs.Article, config proxyStructs.AppConfig, logs *log.Logger) db {
	var myDb db
	var err error
	if config.BDType == "postgresDB"{
		myDb, err = postgresDB.New(ch, config, logs)
		if err != nil{
			logs.Println(err.Error())
			logs.Println("using memoryDB...")
			myDb = MemoryDB.New(ch, config, logs)
		}
	}
	return myDb
}

