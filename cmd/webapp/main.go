package main

import (
	"flag"
	"log"
	"net/http"
	"rvweb/web"

	"github.com/jmoiron/sqlx"
)

var port = flag.String("port", "8080", "port to serve on")

var exampleSchema = `
    CREATE TABLE IF NOT EXISTS posts (id INTEGER PRIMARY KEY, title TEXT, body TEXT)
`
var examplePost = `INSERT INTO posts (title, body) VALUES ('foo', 'bar')`

func main() {
	db, err := sqlx.Connect("sqlite3", ":memory:")
	if err != nil {
		log.Fatalln(err)
	}

	db.MustExec(exampleSchema)
	db.MustExec(examplePost)

	hs := web.NewHttpServer(db)

	log.Println("Starting server on http://localhost:" + *port)
	http.ListenAndServe(":"+*port, hs.Server)
}
