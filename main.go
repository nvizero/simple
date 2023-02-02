package main

import (
	"database/sql"
	"log"
	"simple/api"
	db "simple/db/sqlc"

	_ "github.com/lib/pq"
	"github.com/techschool/simplebank/util"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:root@localhost:5432/simple?sslmode=disable"
	serverAddress = "0.0.0.0:8081"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start("0.0.0.0:8081")
	if err != nil {
		log.Fatal("cannot start server :", err)
	}
}
