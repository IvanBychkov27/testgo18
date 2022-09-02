package main

import (
	"log"
	"net/http"
	"os"
)

type Application struct {
}

func New() *Application {
	return &Application{}
}

func (app *Application) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	res := req.RemoteAddr
	_, err := rw.Write([]byte(res))
	if err != nil {
		log.Println("error write ", err.Error())
	}
	log.Println("call ", req.RemoteAddr)
}

func main() {
	app := New()
	addr := ":80"
	log.Printf("start server - listen %s", addr)
	err := http.ListenAndServe(addr, app)
	if err != nil {
		log.Printf("error %v", err)
		os.Exit(1)
	}
}
