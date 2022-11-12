package entity

type DataQuery struct {
	TABLES  string
	SELECTS string
	WHERE   string
	WHEREOR string
	LIMIT   string
	ORDERBY string
	OFFSET  string
	GROUPBY string
}

const NullString = ""
