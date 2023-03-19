package rssReader

import (
	"encoding/xml"
	"io"
	"log"
	"net/http"
	"newsApp/packages/proxyStructs"
	"sync"
	"time"
)

type reader struct{
	l *log.Logger
	pulse int
	sources []string
	muData sync.Mutex
	data []proxyStructs.Feed
	muNews sync.Mutex
	news []bool
	err error
	ch chan proxyStructs.Article
}

// Создание структуры опрашивающей RSS ленты и пишущей результаты в  общий канал (требуется запуск)
func New(config proxyStructs.AppConfig, logger *log.Logger) *reader {
	return &reader{
		l : logger,
		pulse : config.ReqPer,
		sources : config.Sources,
		muData : sync.Mutex{},
		data : make([]proxyStructs.Feed, len(config.Sources)),
		muNews : sync.Mutex{},
		news : make([]bool, len(config.Sources)),
		ch : make(chan proxyStructs.Article, len(config.Sources) * 25),  // 25 Статей на 1 ссылку!!!
	}
}

// В соответсвии с периодичностью опрсоса определённой в структуре конфигурации
// RSS reader запускает цикл сбора данных и записи в выходной поток
func (r *reader) Start(){
	t := time.NewTicker(time.Duration(r.pulse) * time.Minute)
	go func(){
		for {
			<- t.C
			r.update(getData)
			r.flush()
		}
	}()
	r.l.Println("rss reader started")
}



func (r *reader) update(gD func(string)([]byte, error)){ // обновляет хранимые в reader.data структуры,
															// представляющие xml документы целиком (и их статусы)
	wg := sync.WaitGroup{}
	for n, url := range r.sources{
		wg.Add(1)

		go func(n int, url string,r *reader, wg *sync.WaitGroup){
			defer func(){
				wg.Done()

			}()

			data, err := gD(url)
			if err != nil{
				r.updtErr(err)
				return
			}


			var feed proxyStructs.Feed
			err = xml.Unmarshal(data, &feed)
			if err != nil{
				r.updtErr(err)
				r.l.Println("error via unmarshalling")
				return
			}
			r.muData.Lock()
			r.data[n] = feed
			// r.l.Printf("func update saved: %v articles from source: %v\n", len(feed.Chanel.List), n)
			r.muData.Unlock()
			
		}(n, url, r, &wg)
	}
	wg.Wait()
	if r.err != nil{
		r.l.Println("some errors in reading rss..")
	}
}

func (r *reader) flush(){
	archive := map[int]int{}
	for nS, feed := range r.data{		// Бежим по feed ам - целым xml документам
		archive[nS] = 0
		for _ , article := range feed.Chanel.List{ // Бежим по статьям в Channel каждого feed:
			r.ch <- article                        // пишем в выходной канал
			archive[nS] = archive[nS] + 1
		}

	}
	// r.l.Printf("func flush push to output [source number : articles quantity, ..]: %v\n", archive)
	r.data = make([]proxyStructs.Feed, len(r.sources))
}



func getData(url string) ([]byte, error){
	resp, err := http.Get(url)
	if err != nil{
		return []byte{}, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil{
		return []byte{}, err
	}

	return body, nil
}


func (r *reader) Chan () <- chan proxyStructs.Article{
	return r.ch
}


func (r *reader) updtErr(err error){
	if r.err == nil{
		r.err = err
	}
}

func (r * reader) Err()error{
	return r.err
}
