package main

import (
	"log"
	"net/http"

	"github.com/Hackail/rest-api-go/cmd/server"
)

func main() {
	application := server.New()
	err := http.ListenAndServe(":8080", application)
	if err != nil {
		log.Panic("error initializing server")
	}
}
