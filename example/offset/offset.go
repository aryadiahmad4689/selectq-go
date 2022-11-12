package offset

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
	selectQ := selectq.Init(context.Background())
	db := connectDb()

	selectQ.NewConnect(db).SetTable("test")
	r, _ := selectQ.Read.Select("id").Offset(10).Get()

	fmt.Println(r)
}
