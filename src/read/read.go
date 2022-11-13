package read

import (
	"context"
	"database/sql"
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
	for rows.Next() {
		values := make([]interface{}, len(columns))
		pointers := make([]interface{}, len(columns))
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
	if grb == entity.NullString {
		p.data.GROUPBY = grb
	} else {
		p.data.GROUPBY = "group by " + grb
	}

	return p
}

func (p *Read) Limit(limit int) *Read {
	if limit == 0 {
		p.data.LIMIT = entity.NullString
	} else {
		p.data.LIMIT = "LIMIT " + strconv.Itoa(limit)
	}

	return p
}

func (p *Read) Offset(offset int) *Read {
	if offset == 0 {
		p.data.OFFSET = entity.NullString
	} else {
		p.data.OFFSET = "OFFSET " + strconv.Itoa(offset)
	}

	return p
}

func (p *Read) OrderBy(orderBy string) *Read {
	if orderBy == entity.NullString {
		p.data.ORDERBY = entity.NullString
	} else {
		p.data.ORDERBY = "ORDER BY " + orderBy
	}

	return p
}
