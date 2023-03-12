package filesIO

import (
	"encoding/json"
	"fmt"
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
	fmt.Println(Pwd)
	configFile, err := os.Open(Pwd +"\\NewsApp\\packages\\filesIO\\config.txt")
	if err != nil{
		log.Fatalf("cant open config.txt in files: %v\n", err)
	}

	bytes, err :=  ioutil.ReadAll(configFile)
	err = json.Unmarshal(bytes, &AConf)
	if err != nil{
		log.Fatalf("cant read config.txt: %v\n", err)
	}
}
