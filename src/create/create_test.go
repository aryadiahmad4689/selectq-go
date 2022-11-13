package create

import (
	"reflect"
	"testing"
)

func TestCreate_SetTable(t *testing.T) {
	type args struct {
		table string
	}
	tests := []struct {
		name string
		c    *Create
		args args
		want *Create
	}{
		{
			name: "testing set table",
			c:    &Create{},
			args: args{
				table: "test",
			},
			want: &Create{
				table: "test",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.SetTable(tt.args.table); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create.SetTable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreate_AddField(t *testing.T) {
	type args struct {
		colums string
		value  interface{}
	}
	tests := []struct {
		name string
		c    *Create
		args args
		want *Create
	}{
		// TODO: Add test cases.
		{
			name: "testing set table",
			c:    &Create{},
			args: args{
				colums: "id",
				value:  "1",
			},
			want: &Create{
				column:     []string{"id"},
				value:      []interface{}{"1"},
				colNumber:  []string{"$1"},
				counterCol: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.AddField(tt.args.colums, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create.AddField() = %v, want %v", got, tt.want)
			}
		})
	}
}
