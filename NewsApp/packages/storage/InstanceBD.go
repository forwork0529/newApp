package storage

import (
	"log"
	"newsApp/packages/proxyStructs"
	"newsApp/packages/storage/MemoryDB"
	"newsApp/packages/storage/postgresDB"
)

type DB interface{
	Push(art proxyStructs.Article)
	Get(number int)([]proxyStructs.Article, error)
	Start()
	Flush()
}

func New(ch <- chan proxyStructs.Article, config proxyStructs.AppConfig, logs *log.Logger) DB {
	var myDb DB
	var err error
	if config.BDType == "postgresDB"{
		myDb, err = postgresDB.New(ch, config, logs)
		if err != nil{
			logs.Println(err.Error())

			myDb = MemoryDB.New(ch, config, logs)
		}
	}

	return myDb
}

