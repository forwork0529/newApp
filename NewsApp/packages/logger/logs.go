package logger

import (
	"fmt"
	"log"
	"newsApp/packages/filesIO"
	"os"
)

func New()*log.Logger{
	o, err := os.Create(filesIO.Pwd +"\\NewsApp\\packages\\filesIO\\output.txt")
	fmt.Println()
	if err != nil{
		log.Fatal(err)
	}
	return log.New(o, "INFO: ", log.LstdFlags )
}
