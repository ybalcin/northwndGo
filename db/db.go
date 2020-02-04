package db

import (
	"context"
	"database/sql"
	_ "github.com/denisenkom/go-mssqldb"
	"log"
)

func ConnectDb() (*sql.DB, context.Context) {
	connectionString := "server=localhost;database=NORTHWND;user id=northwndUser;password=r}-c(%3P8P"

	// Create connection pool
	db, err := sql.Open("sqlserver", connectionString)
	if err != nil {
		log.Fatal("Error:", err.Error())
	}
	ctx := context.Background()

	return db, ctx
}
