package sqlc

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal(err)
	}

	driver := "postgres"
	src := os.Getenv("DB_URL")

	conn, err := sql.Open(driver, src)
	if err != nil {
		log.Fatal(err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
