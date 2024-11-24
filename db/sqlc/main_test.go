package db

import (
	"context"
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver string = "postgres"
	dbSource string = "postgresql://ROOT:secret@localhost:5432/mysimple-bank?sslmode=disable"
)

var testQueries *Queries
var conn *sql.DB
var ctx = context.Background()

func TestMain(m *testing.M) {
	var err error
	conn, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatalln("can't connect to db:", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
