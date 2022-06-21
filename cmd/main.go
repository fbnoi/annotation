package main

import (
	"os"

	"fbnoi.com/annotation/direct"
)

func main() {
	router := &direct.Route{
		Name:       "Index",
		Path:       "/",
		HandleFunc: "main.Index",
	}
	router.Render(os.Stdout)
}
