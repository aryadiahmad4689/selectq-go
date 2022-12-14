package read

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"

	"github.com/aryadiahmad4689/selectq-go/src/entity"
)

type Read struct {
	db         *sql.DB
	table      string
	query      string
	where      []string
	whereOr    []string
	searchWere []interface{}
	data       entity.DataQuery
}

func Init(ctx context.Context, conn *sql.DB) *Read {
	return &Read{
		db: conn,
	}
}

func (p *Read) WhereOr(where string, search string) *Read {
	p.whereOr = append(p.whereOr, where)
	p.searchWere = append(p.searchWere, search)
	return p
}

func (p *Read) Get() ([]map[string]interface{}, error) {
	query := p.generateQuery()
	var objects = []map[string]interface{}{}

	rows, err := p.db.QueryContext(context.Background(), query, p.searchWere...)
	if err != nil {
		return objects, err
	}
	columns, err := rows.Columns()
	if err != nil {
		return objects, err
	}

	values := make([]interface{}, len(columns))
	pointers := make([]interface{}, len(columns))
	for rows.Next() {

		for i, _ := range values {
			pointers[i] = &values[i]
		}
		rows.Scan(pointers...)
		resultMap := make(map[string]interface{})
		for i, val := range values {
			resultMap[columns[i]] = val
		}
		objects = append(objects, resultMap)
	}
	return objects, nil
}

func (p *Read) Select(selects string) *Read {
	p.data.SELECTS = selects
	return p
}

func (p *Read) Where(where string, find string) *Read {
	p.where = append(p.where, where)
	p.searchWere = append(p.searchWere, find)
	return p
}

func (p *Read) GroupBy(grb string) *Read {
	if grb != entity.NullString {
		p.data.GROUPBY = entity.GroupBy + grb
	}

	return p
}

func (p *Read) Limit(limit int) *Read {
	if limit != 0 {
		p.data.LIMIT = entity.Limit + strconv.Itoa(limit)
	}

	return p
}

func (p *Read) Offset(offset int) *Read {
	if offset != 0 {
		p.data.OFFSET = entity.Offset + strconv.Itoa(offset)
	}

	return p
}

func (p *Read) OrderBy(orderBy string) *Read {
	if orderBy != entity.NullString {
		p.data.ORDERBY = entity.OrderBy + orderBy
	}

	return p
}

func (p *Read) InnerJoin(table string, column string, coloumJoin string) *Read {
	if table != entity.NullString || column != entity.NullString || coloumJoin != entity.NullString {
		p.data.JOIN += entity.InnerJOIN + table + entity.On + fmt.Sprintf("%s.%s ", table, column) + entity.Equal + coloumJoin + entity.SpaceEmptyString
	}
	return p
}

func (p *Read) LeftJoin(table string, column string, coloumJoin string) *Read {
	if table != entity.NullString || column != entity.NullString || coloumJoin != entity.NullString {
		p.data.JOIN += entity.LeftJoin + table + entity.On + fmt.Sprintf("%s.%s ", table, column) + entity.Equal + coloumJoin + entity.SpaceEmptyString
	}
	return p
}

func (p *Read) RightJoin(table string, column string, coloumJoin string) *Read {
	if table != entity.NullString || column != entity.NullString || coloumJoin != entity.NullString {
		p.data.JOIN += entity.RightJoin + table + entity.On + fmt.Sprintf("%s.%s ", table, column) + entity.Equal + coloumJoin + entity.SpaceEmptyString
	}
	return p
}
