package selectq

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/lib/pq"
)

func connectDb() *sql.DB {

	db, err := sql.Open("postgres", "postgresql://postgres:postgres@localhost:5430/postgres?sslmode=disable")
	if err != nil {
		panic(err)
	}

	// defer db.Close()
	if err = db.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("success connect")

	return db
}

func TestApp(t *testing.T) {

}
