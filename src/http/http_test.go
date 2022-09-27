package http

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestHttp(t *testing.T) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world!")
	})
	http.HandleFunc("/time/", func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		timeStr := fmt.Sprintf(`{"time":"%s"}`, t)
		w.Write([]byte(timeStr))
	})
	http.ListenAndServe("localhost:8080", nil)
}
