package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/badis/hackathon/internal/handler"
	"github.com/badis/hackathon/internal/service"
	"github.com/hako/branca"

	_ "github.com/jackc/pgx/stdlib"
)

const (
	databaseURL = "postgresql://admin:0000@database:5432/hackathon_db?sslmode=disable"
	port        = 5000
)

func main() {
	db, err := sql.Open("pgx", databaseURL)
	if err != nil {
		log.Fatalf("could not open db connection: %v\n", err)
		return
	}

	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatalf("could not ping to db : %v\n", err)
		return
	}

	codec := branca.NewBranca("supersecretkeysupersecretkeysupe")

	s := service.New(db, codec)

	h := handler.New(s)

	log.Printf("accepting connections on port: %d\n", port)

	if err = http.ListenAndServe(fmt.Sprintf(":%d", port), h); err != nil {
		log.Fatalf("could not start server: %v\n", err)
	}
}
