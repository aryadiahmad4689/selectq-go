package selectq

import (
	"context"
	"database/sql"

	"github.com/aryadiahmad4689/selectq-go/src/create"
	"github.com/aryadiahmad4689/selectq-go/src/read"
)

type SelectQ struct {
	Read   *read.Read
	Create *create.Create
}

func Init(ctx context.Context, conn *sql.DB) *SelectQ {
	return &SelectQ{
		Read:   read.Init(ctx, conn),
		Create: create.Init(ctx, conn),
	}
}
