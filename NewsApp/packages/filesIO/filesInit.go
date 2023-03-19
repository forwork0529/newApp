package filesIO

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"newsApp/packages/proxyStructs"
	"os"
)

var (
	Pwd   string
	err   error
	AConf proxyStructs.AppConfig
)



func init(){
	Pwd, err = os.Getwd()
	if err != nil{
		log.Fatal(err)
	}
	var bytes []byte
	configFile, err := os.Open(Pwd +"\\NewsApp\\packages\\filesIO\\config.json")

	if err != nil{
		log.Printf("cant open config.json in files: %v\n", err)
		log.Print("use alternative config file..")
		bytes = []byte("{\n   \"rss\":[\n        \"https://habr.com/ru/rss/hub/go/all/?fl=ru\",\n        \"https://habr.com/ru/rss/best/daily/?fl=ru\",\n        \"https://cprss.s3.amazonaws.com/golangweekly.com.xml\"\n   ],\n   \"request_period\": 1,\n   \"dbType\":\"postgresDB\",\n   \"connString\":\"postgres://postgres:postgres@127.0.0.1:8081/articles\"\n}")
	}else{
		bytes, err =  ioutil.ReadAll(configFile)
	}
	err = json.Unmarshal(bytes, &AConf)
	if err != nil{
		log.Fatalf("cant read config.json: %v\n", err)
	}


}
