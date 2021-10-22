package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/badis/hackathon/internal/handler"
	"github.com/badis/hackathon/internal/service"
	"github.com/hako/branca"

	_ "github.com/jackc/pgx/stdlib"
)

var (
	dbuser      = os.Getenv("POSTGRES_USER")
	dbpassword  = os.Getenv("POSTGRES_PASSWORD")
	dbname      = os.Getenv("POSTGRES_DB")
	dbhost      = os.Getenv("POSTGRES_ADDR")
	databaseURL = fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=disable", dbuser, dbpassword, dbhost, dbname)
	port        = os.Getenv("PORT")
)

func main() {
	db, err := sql.Open("pgx", databaseURL)
	if err != nil {
		log.Fatalf("could not open db connection: %v\n", err)
		return
	}

	defer db.Close()

	/* if err = db.Ping(); err != nil {
		log.Fatalf("could not ping to db : %v\n", err)
		// return
	}*/

	codec := branca.NewBranca("supersecretkeysupersecretkeysupe")

	s := service.New(db, codec)

	h := handler.New(s)

	log.Printf("accepting connections on port: %s\n", port)

	if err = http.ListenAndServe(fmt.Sprintf("%s:%s", "0.0.0.0", port), h); err != nil {
		log.Fatalf("could not start server: %v\n", err)
	}
}
