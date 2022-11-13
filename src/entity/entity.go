package entity

type DataQuery struct {
	TABLES    string
	SELECTS   string
	WHERE     string
	WHEREOR   string
	LIMIT     string
	ORDERBY   string
	OFFSET    string
	GROUPBY   string
	JOIN      string
	COLUMN    string
	COLNUMBER string
}

const (
	NullString       = ""
	Limit            = "LIMIT "
	Offset           = "OFFSET "
	OrderBy          = "ORDER BY "
	GroupBy          = "GROUP BY "
	InnerJOIN        = "INNER JOIN "
	LeftJoin         = "LEFT JOIN "
	RightJoin        = "RIGHT JOIN "
	On               = " ON "
	Equal            = "= "
	SpaceEmptyString = " "
	DollarQuote      = "$"
	OneNumber        = 1
)
