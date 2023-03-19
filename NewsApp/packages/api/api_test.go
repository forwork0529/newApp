package api

import (
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"newsApp/packages/proxyStructs"
	"newsApp/packages/storage/MemoryDB"
	"testing"
)

func TestAPI(t *testing.T) {

	logs := log.Default()
	logs.Println()
	config := proxyStructs.AppConfig{}
	ch := make(<-chan proxyStructs.Article, 1)
	bd := MemoryDB.New(ch, config, logs)
	bd.Get(5)
	myAPI := New(bd, logs)
	handler := http.Handler(myAPI.Router())
	req := httptest.NewRequest(http.MethodGet, "/news/1", nil)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Fatalf("got status code: %v, want: %v\n", rr.Code, http.StatusOK)
	}
	b, err := io.ReadAll(rr.Body)
	if err != nil {
		t.Fatal(err)
	}
	want := "[{\"ID\":123,\"Title\":\"Attendant article\",\"Content\":\"Some description it the body of article\"," +
		"\"PubTime\":123456,\"Link\":\"https://golangweekly.com/issues/451\"}]\n"

	if string(b) != want{
		t.Errorf("got: %v /// != want: %v ///\n", len(string(b)), len(want) )
	}

}