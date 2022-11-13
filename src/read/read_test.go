package read

import (
	"reflect"
	"testing"

	"github.com/aryadiahmad4689/selectq-go/src/entity"
)

func TestRead_WhereOr(t *testing.T) {
	type args struct {
		where  string
		search string
	}
	tests := []struct {
		name string
		args args
		want *Read
	}{
		{
			name: "test where",
			args: args{
				search: "halo",
				where:  "id =?",
			},
			want: &Read{
				whereOr:    []string{"id =?"},
				searchWere: []interface{}{"halo"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Read{}
			if got := p.WhereOr(tt.args.where, tt.args.search); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Read.WhereOr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRead_Select(t *testing.T) {
	type args struct {
		selects string
	}
	tests := []struct {
		name string
		p    *Read
		args args
		want *Read
	}{
		// TODO: Add test cases.
		{
			name: "test where",
			p:    &Read{},
			args: args{
				selects: "username",
			},
			want: &Read{
				data: entity.DataQuery{
					SELECTS: "username",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Select(tt.args.selects); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Read.Select() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRead_Where(t *testing.T) {
	type args struct {
		where string
		find  string
	}
	tests := []struct {
		name string
		args args
		want *Read
	}{
		// TODO: Add test cases.
		{
			name: "test where",
			args: args{
				where: "id =?",
				find:  "1",
			},
			want: &Read{
				where:      []string{"id =?"},
				searchWere: []interface{}{"1"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Read{}
			if got := p.Where(tt.args.where, tt.args.find); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Read.Where() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRead_GroupBy(t *testing.T) {
	type args struct {
		grb string
	}
	tests := []struct {
		name string
		p    *Read
		args args
		want *Read
	}{
		// TODO: Add test cases.
		{
			name: "test where",
			args: args{
				grb: "id",
			},
			want: &Read{
				data: entity.DataQuery{
					GROUPBY: "GROUP BY id",
				},
			},
			p: &Read{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.GroupBy(tt.args.grb); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Read.GroupBy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRead_Limit(t *testing.T) {
	type args struct {
		limit int
	}
	tests := []struct {
		name string
		p    *Read
		args args
		want *Read
	}{
		// TODO: Add test cases.
		{
			name: "test where",
			args: args{
				limit: 10,
			},
			want: &Read{
				data: entity.DataQuery{
					LIMIT: "LIMIT 10",
				},
			},
			p: &Read{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Limit(tt.args.limit); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Read.Limit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRead_Offset(t *testing.T) {
	type args struct {
		offset int
	}
	tests := []struct {
		name string
		p    *Read
		args args
		want *Read
	}{
		// TODO: Add test cases.
		{
			name: "test where",
			args: args{
				offset: 10,
			},
			want: &Read{
				data: entity.DataQuery{
					OFFSET: "OFFSET 10",
				},
			},
			p: &Read{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Offset(tt.args.offset); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Read.Offset() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRead_OrderBy(t *testing.T) {
	type args struct {
		orderBy string
	}
	tests := []struct {
		name string
		p    *Read
		args args
		want *Read
	}{
		// TODO: Add test cases.
		{
			name: "test where",
			args: args{
				orderBy: "username asc",
			},
			want: &Read{
				data: entity.DataQuery{
					ORDERBY: "ORDER BY username asc",
				},
			},
			p: &Read{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.OrderBy(tt.args.orderBy); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Read.OrderBy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRead_InnerJoin(t *testing.T) {
	type args struct {
		table      string
		column     string
		coloumJoin string
	}
	tests := []struct {
		name string
		p    *Read
		args args
		want *Read
	}{
		// TODO: Add test cases.
		{
			name: "test where",
			args: args{
				table:      "test",
				column:     "id",
				coloumJoin: "test2.id",
			},
			want: &Read{
				data: entity.DataQuery{
					JOIN: "INNER JOIN test ON test.id = test2.id ",
				},
			},
			p: &Read{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.InnerJoin(tt.args.table, tt.args.column, tt.args.coloumJoin); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Read.InnerJoin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRead_LeftJoin(t *testing.T) {
	type args struct {
		table      string
		column     string
		coloumJoin string
	}
	tests := []struct {
		name string
		p    *Read
		args args
		want *Read
	}{
		// TODO: Add test cases.
		{
			name: "test where",
			args: args{
				table:      "test",
				column:     "id",
				coloumJoin: "test2.id",
			},
			want: &Read{
				data: entity.DataQuery{
					JOIN: "LEFT JOIN test ON test.id = test2.id ",
				},
			},
			p: &Read{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.LeftJoin(tt.args.table, tt.args.column, tt.args.coloumJoin); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Read.LeftJoin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRead_RightJoin(t *testing.T) {
	type args struct {
		table      string
		column     string
		coloumJoin string
	}
	tests := []struct {
		name string
		p    *Read
		args args
		want *Read
	}{
		// TODO: Add test cases.
		{
			name: "test where",
			args: args{
				table:      "test",
				column:     "id",
				coloumJoin: "test2.id",
			},
			want: &Read{
				data: entity.DataQuery{
					JOIN: "RIGHT JOIN test ON test.id = test2.id ",
				},
			},
			p: &Read{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.RightJoin(tt.args.table, tt.args.column, tt.args.coloumJoin); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Read.RightJoin() = %v, want %v", got, tt.want)
			}
		})
	}
}
