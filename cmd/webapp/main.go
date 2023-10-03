package main

import (
	"log"
	"net/http"

	"rvweb"
)

const PORT = ":8080"

func main() {
	log.Println("Starting server on http://localhost" + PORT)
	hs := rvweb.NewHttpServer()
	http.ListenAndServe(PORT, hs.Server)
}
