package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/livingdolls/golang-transfer-money/api"
	db "github.com/livingdolls/golang-transfer-money/db/sqlc"
	"github.com/livingdolls/golang-transfer-money/util"
)

func main() {
	config, err := util.LoadConfig(".")
	conn, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("cannot connect to db :", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)

	if err != nil {
		log.Fatal("cannot create server: ", err)
	}

	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatal("cannot start server", err)
	}
}
