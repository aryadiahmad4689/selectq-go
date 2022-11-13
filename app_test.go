package selectq

import (
	"context"
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
	db := connectDb()
	selectQ := Init(context.Background(), db)
	selectQ.Read.SetTable("order_header")

	r, err := selectQ.Read.Select("chanel").LeftJoin("order_detail", "id", "order_header.id").Where("order_header.id =$1", "4").WhereOr("order_header.id !=$2", "10").GroupBy("order_header.id").Limit(10).OrderBy("order_header.id asc").Get()

	fmt.Println(r, err)
}
