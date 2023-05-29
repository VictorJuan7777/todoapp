package main

import (
	"database/sql"
	"log"
	"todoapp/api"
	db "todoapp/db/sqlc"

	_ "github.com/lib/pq"
)

func main() {
	DBDriver := "postgres"
	DBSource := "postgresql://root:secret@10.53.128.3:5432/todoapp?sslmode=disable"
	// DBSource := "postgresql://root:secret@35.229.174.192:5432/todoapp?sslmode=disable"
	conn, err := sql.Open(DBDriver, DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", DBDriver, err)
	}
	store := db.New(conn)
	//store := db.NewStore(conn)
	//gin.SetMode(gin.ReleaseMode)
	err = api.NewGinServer(store)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}

}
