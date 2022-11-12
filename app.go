package selectq

import (
	"context"
	"database/sql"

	"github.com/aryadiahmad4689/selectq-go/src/read"
)

type SelectQ struct {
	Read *read.Read
}

func Init(ctx context.Context) *SelectQ {
	return &SelectQ{
		Read: read.Init(ctx),
	}
}

func (s *SelectQ) SetTable(table string) *SelectQ {
	s.Read.Table = table
	return s
}

func (s *SelectQ) NewConnect(db *sql.DB) *SelectQ {
	s.Read.DB = db
	return s
}
