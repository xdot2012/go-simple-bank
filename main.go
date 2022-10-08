package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/xdot2012/simple-bank/api"
	db "github.com/xdot2012/simple-bank/db/sqlc"
	"github.com/xdot2012/simple-bank/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot Load Configuration:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot Connect to Database", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Could not Start Server")
	}
}
