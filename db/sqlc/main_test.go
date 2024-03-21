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

func TestMain(t *testing.M) {
	db, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	testQueries = New(db)

	os.Exit(t.Run())
}
