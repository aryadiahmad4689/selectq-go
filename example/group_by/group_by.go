package group_by

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/aryadiahmad4689/selectq-go"
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

func main() {
	db := connectDb()
	selectQ := selectq.Init(context.Background(), db)

	selectQ.Read.SetTable("test")

	r, _ := selectQ.Read.Select("id").GroupBy("id").Get()

	fmt.Println(r)
}
