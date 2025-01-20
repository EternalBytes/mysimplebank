package db

import (
	"context"
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/eternalbytes/simplebank/util"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var conn *sql.DB
var ctx = context.Background()

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../../")
	if err != nil {
		log.Fatalln("cannot load config:", err)
	}
	conn, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatalln("can't connect to db:", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
