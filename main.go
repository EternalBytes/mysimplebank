package main

import (
	"database/sql"
	"log"

	"github.com/eternalbytes/simplebank/api"
	db "github.com/eternalbytes/simplebank/db/sqlc"
	"github.com/eternalbytes/simplebank/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatalln("cannot load cofigurations", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatalln("can't connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatalln("cannot start server", err)
	}
}
