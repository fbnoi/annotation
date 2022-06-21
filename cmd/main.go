package main

import (
	"io"
	"net/http"
	"os"

	"fbnoi.com/annotation/direct"
)

func main() {
	router := &direct.Route{
		Name:       "Index",
		Path:       "/",
		HandleFunc: "main.Index",
		Methods:    []string{"GET", "POST"},
	}
	err := direct.Render(os.Stdout, router)
	if err != nil {
		panic(err)
	}
}

//@Route(/post/:id, name=post_show, methods=[GET])
func Show(req *http.Request, w http.ResponseWriter) {
	io.WriteString(w, "hello world")
}

//@Route(/post/:id, name=post_edit, methods=[POST])
func Edit(req *http.Request, w http.ResponseWriter) {
	io.WriteString(w, "hello world")
}

//@Filter({paths:[/post/*], order:1})
func MD1(req *http.Request, w http.ResponseWriter) {
}

//@Filter({paths: [/], order: 1})
func MD2(req *http.Request, w http.ResponseWriter) {
}

//@Filter({paths: [/], order: 1})
func MD3(req *http.Request, w http.ResponseWriter) {
}
