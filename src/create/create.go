package create

import (
	"context"
	"database/sql"

	"github.com/aryadiahmad4689/selectq-go/src/entity"
)

type Create struct {
	db         *sql.DB
	table      string
	query      string
	column     []string
	value      []interface{}
	colNumber  []string
	counterCol int
	data       entity.DataQuery
}

func Init(ctx context.Context, conn *sql.DB) *Create {
	return &Create{
		db: conn,
	}
}

func (c *Create) SetTable(table string) *Create {
	c.table = table
	return c
}

func (c *Create) AddField(colums string, value interface{}) *Create {
	c.column = append(c.column, colums)
	c.value = append(c.value, value)
	val := c.getColNumber()
	c.colNumber = append(c.colNumber, val)
	return c
}

func (c *Create) Save(ctx context.Context) (string, error) {
	query := c.generateQuery()
	var id string
	err := c.db.QueryRowContext(ctx, query, c.value...).Scan(&id)
	return id, err
}
