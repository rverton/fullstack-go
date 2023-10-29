package main

import (
	"flag"
	"log"
	"net/http"
	"rvweb/repository"
	"rvweb/web"

	"github.com/jmoiron/sqlx"
)

var port = flag.String("port", "8080", "port to serve on")

func main() {

	// initial DB connection
	db, err := sqlx.Connect("sqlite3", ":memory:")
	if err != nil {
		log.Fatalln(err)
	}

	// run migrations
	err = repository.Migrate(db.DB, "file://migrations")
	if err != nil {
		log.Fatalln(err)
	}

	// initialize http server and pass DB
	hs := web.NewHttpServer(db)

	log.Println("Starting server on http://localhost:" + *port)
	http.ListenAndServe(":"+*port, hs.Server)
}
