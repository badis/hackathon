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

func main() {

	var (
		dbUser     = os.Getenv("DB_USER")
		dbPassword = os.Getenv("DB_PASSWORD")
		dbHost     = os.Getenv("DB_HOST")
		dbPort     = os.Getenv("DB_PORT")
		dbName     = os.Getenv("DB_NAME")

		databaseURL = fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)
		port        = os.Getenv("API_PORT")
	)

	log.Printf(databaseURL)

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

	log.Printf("accepting connections on port: %s\n", port)

	if err = http.ListenAndServe(fmt.Sprintf(":%s", port), h); err != nil {
		log.Fatalf("could not start server: %v\n", err)
	}
}
