package rssReader

import (
	"log"
	"newsApp/packages/proxyStructs"
	"sync"
	"testing"
	"time"
)

func getDataTest(url string)([]byte, error){
	return []byte(Data1), nil
}

func readFromChan(ch <-chan proxyStructs.Article,logs *log.Logger)(res bool){
	res = true
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func(wg *sync.WaitGroup, logs *log.Logger, res *bool){
		select{
		case <- ch: // a := <- ch:
			//logs.Printf("fetched title from struct which was read from chan: %v\n", a.Title)
		case <- time.After(3 * time.Second)	:
			*res = false
			logs.Println("time for reading from channel elapsed")
		}
		wg.Done()
	}(&wg, logs, &res)
	wg.Wait()
	return res
}



func Test_RssReader(t *testing.T) {

	config := proxyStructs.AppConfig{
			Sources: []string {"https://helloWorld.ru"},
			ReqPer: 10,
	}

	r := New(config, log.Default())
	r.l.Println("reader created")
	t.Run("update testing", func (t *testing.T){
		r.update(getDataTest)
		if err := r.Err(); err != nil{
			t.Fatalf("error via updating: %v", r.Err().Error())
		}

		if len(r.data[0].Chanel.List[0].Description) < len("There is a long article!"){
			t.Fatal("wrong data in description saved")
		}
		r.l.Printf("fetched title from xml: %v\n", r.data[0].Chanel.List[0].Title )

		r.flush()

		output := r.Chan()
		if !readFromChan(output, r.l){
			t.Fatal("reading from channel fallen")
		}

	})
}




func Test_RssReaderNet (t *testing.T){
	config := proxyStructs.AppConfig{
		Sources: []string {"https://habr.com/ru/rss/hub/go/all/?fl=ru"},
		ReqPer: 10,
	}
	r := New(config, log.Default())
	r.l.Println("reader created")
	t.Run("update from NET testing", func (t *testing.T){
		r.update(getData)
		if err := r.Err(); err != nil{
			t.Fatalf("error via updating: %v", r.Err().Error())
		}

		if len(r.data[0].Chanel.List[0].Description) < len("There is a long article!"){
			t.Fatal("wrong data in description saved")
		}
		// r.l.Printf("fetched title from saved by reader data : %v\n", r.data[0].Chanel.List[0].Title )

		nOfSavedArticles := len(r.data[0].Chanel.List)

		r.l.Printf("there are %v articles in the RSSreader\n", len(r.data[0].Chanel.List))

		r.flush()

		output := r.Chan()

		for n := 0 ;n < nOfSavedArticles; n ++{
			if !readFromChan(output, r.l){
				t.Fatalf("reading from channel fallen, article number %v of total number %v\n", n, nOfSavedArticles)
			}
		}
	})

}

