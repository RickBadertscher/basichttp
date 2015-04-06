package main

import (
	"fmt"
	//	"log"
	"net/http"
)

type DefaultHandler struct {
	hits uint64
}

type handlerMap map[string]func(http.ResponseWriter, *http.Request)

var mux handlerMap

func main() {
	server := http.Server{
		Addr:    ":8080",
		Handler: &DefaultHandler{},
	}
	mux = make(handlerMap)
	mux["/"] = root
	//	mux["/favicon.ico"] = iconHandler

	server.ListenAndServe()
}

func (o *DefaultHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	if handler, ok := mux[req.URL.String()]; ok {
		o.hits += 1
		fmt.Fprintf(res, "The router, hits = %d\n", o.hits)
		fmt.Printf("req = %v\n", req)
		handler(res, req)
		return
	}

	http.NotFound(res, req)
}

func root(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "this is the context root\n")
}
