package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/georgifotev1/go-api/aplication"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	connStr := os.Getenv("DB_URL")
	if connStr == "" {
		log.Fatal("DB_URL not foun in the enviroment")
	}

	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(fmt.Errorf("can't connect to database: %v", err))
	}

	if err := conn.Ping(); err != nil {
		log.Fatal(fmt.Errorf("connection not established: %v", err))
	}

	defer func() {
		if err := conn.Close(); err != nil {
			fmt.Println("failed to close database", err)
		}
	}()

	app := aplication.New(conn)

	log.Fatal(app.Start())
}
