package create

import (
	"strconv"
	"strings"

	"github.com/aryadiahmad4689/selectq-go/src/entity"
	"github.com/aryadiahmad4689/selectq-go/src/utils"
)

func (c *Create) getQuery() *Create {
	c.query = `INSERT INTO ` + c.table + " ({{.COLUMN}}) VALUES({{.COLNUMBER}}) RETURNING id"
	return c
}

func (c *Create) generateColumn() string {
	colums := strings.Join(c.column[:], ",")
	if colums != entity.NullString {
		c.data.COLUMN = colums
	}
	return colums
}

func (c *Create) generateColNumber() *Create {
	colNumber := strings.Join(c.colNumber[:], ",")
	if colNumber != entity.NullString {
		c.data.COLNUMBER = colNumber
	}
	return c
}

func (c *Create) generateQuery() string {
	c.getQuery()
	c.generateColumn()
	c.generateColNumber()
	query := utils.Parse("generate", c.query, c.data)
	return query
}

func (c *Create) getColNumber() string {
	c.counterCol += entity.OneNumber
	val := entity.DollarQuote + strconv.Itoa(c.counterCol)
	return val
}
