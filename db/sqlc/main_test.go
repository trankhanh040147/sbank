package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
	"testing"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/sbank?sslmode=disable"
)

// might be db connection or transaction
var testQueries *Queries
var testDB *sql.DB

func TestMain(t *testing.M) {
	var err error

	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	testQueries = New(testDB)

	os.Exit(t.Run())
}
