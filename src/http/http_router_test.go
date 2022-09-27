package http

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"testing"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Welcome!\n")
}
func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "Hello, %s!\n", ps.ByName("name"))
}

func TestHttpRouterPackage(t *testing.T) {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
