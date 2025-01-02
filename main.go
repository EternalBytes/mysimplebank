package main

import (
	"database/sql"
	"log"

	"github.com/eternalbytes/simplebank/api"
	db "github.com/eternalbytes/simplebank/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDriver      string = "postgres"
	dbSource      string = "postgresql://ROOT:secret@localhost:5432/mysimple-bank?sslmode=disable"
	serverAddress string = "localhost:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatalln("can't connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatalln("cannot start server", err)
	}
}
