package read

import (
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

func (p *Read) getQuery() *Read {
	p.query = "SELECT {{.SELECTS}} FROM {{.TABLES}} {{.WHERE}} {{.WHEREOR}} {{.GROUPBY}} {{.ORDERBY}} {{.OFFSET}} {{.LIMIT}}"
	return p
}

func (p *Read) getTable() {
	p.data.TABLES = p.Table
}

func (p *Read) generateQuery() string {
	p.getTable()
	p.generateWhere()
	p.generateWhereOr()
	p.getQuery()

	query := utils.Parse("generate", p.query, p.data)
	return query
}
