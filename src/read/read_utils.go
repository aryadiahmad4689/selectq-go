package read

import (
	"database/sql"
	"strings"

	"github.com/aryadiahmad4689/selectq-go/src/entity"
	"github.com/aryadiahmad4689/selectq-go/src/utils"
)

func (p *Read) generateWhere() *Read {
	data := strings.Join(p.where[:], " and ")
	if data == entity.NullString {
		p.data.WHERE = data
	} else {
		p.data.WHERE = "where " + data

	}
	return p
}

func (p *Read) SetTable(table string) *Read {
	p.table = table
	return p
}

func (p *Read) SetDb(db *sql.DB) *Read {
	p.db = db
	return p
}

func (p *Read) getQuery() *Read {
	p.query = "SELECT {{.SELECTS}} FROM {{.TABLES}} {{.WHERE}} {{.WHEREOR}} {{.GROUPBY}} {{.ORDERBY}} {{.OFFSET}} {{.LIMIT}}"
	return p
}

func (p *Read) generateWhereOr() *Read {
	data := strings.Join(p.whereOr[:], " OR ")
	if data == entity.NullString {
		p.data.WHEREOR = data
	} else if len(p.data.WHERE) > 1 && data != entity.NullString {
		p.data.WHEREOR = " OR " + data
	} else {
		p.data.WHEREOR = "WHERE " + data

	}
	return p
}

func (p *Read) getTable() {
	p.data.TABLES = p.table
}

func (p *Read) generateQuery() string {
	p.getTable()
	p.generateWhere()
	p.generateWhereOr()
	p.getQuery()

	query := utils.Parse("generate", p.query, p.data)
	return query
}
