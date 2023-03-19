package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"newsApp/packages/filesIO"
	"newsApp/packages/proxyStructs"
	"newsApp/packages/storage"
	"strconv"
)

type API struct{
	router * mux.Router
	db storage.DB
	logs *log.Logger
}
// Создаются структура (API), содержащая маршрутизатор запросов и реализацию БД,
// путям в маршрутизаторе  назначаются методы самой структуры API.
func New(db storage.DB,l *log.Logger) *API{
	api := &API{router : mux.NewRouter(), db : db, logs : l}
	api.endpoints()
	api.logs.Println("api created")
	return api
	}

func (a *API) endpoints(){
	a.router.HandleFunc("/news/{n}", a.posts).Methods(http.MethodGet, http.MethodOptions)
	a.router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(filesIO.Pwd + "\\NewsApp\\cmd\\webapp"))))
}

// Метод API возвращающий клиенту публикации из БД в формате json.
func(a *API) posts(w http.ResponseWriter, q *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if q.Method == http.MethodOptions{
		return
	}
	s, _ := mux.Vars(q)["n"]
	n, _ := strconv.Atoi(s)
	list, err := a.db.Get(n)
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)

	}
	news := make([]proxyStructs.Post, len(list))
	for i := range list{
		news[i].ID = list[i].ID
		news[i].Title = list[i].Title
		news[i].Content = list[i].Description
		news[i].PubTime = int64(list[i].PubDate)
		news[i].Link = list[i].Link
	}

	json.NewEncoder(w).Encode(news)
}


// Функция возвращает маршрутизатор для передачи серверу
func (a *API) Router() *mux.Router{
	return a.router
}