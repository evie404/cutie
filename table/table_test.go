package table

import (
	"fmt"
	"math/rand"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTablesFromConfig(t *testing.T) {
	type args struct {
		configFilepath string
	}
	tests := []struct {
		name    string
		args    args
		want    []Table
		wantErr bool
	}{
		{
			"errors on nonexistent file",
			args{
				fmt.Sprintf("%v.yaml", rand.Int63()),
			},
			nil,
			true,
		},
		{
			"errors on malformed file",
			args{
				filepath.Join("testdata", "bad_tables.yaml"),
			},
			nil,
			true,
		},
		{
			"errors on file with invalid tables",
			args{
				filepath.Join("testdata", "invalid_tables.yaml"),
			},
			nil,
			true,
		},
		{
			"load tables1.yaml",
			args{
				filepath.Join("testdata", "tables1.yaml"),
			},
			[]Table{
				{
					TableName:        "table_ones",
					FilenameOverride: "table_one",
					ClassName:        "TableOne",
				},
				{
					TableName:        "table_twos",
					FilenameOverride: "table_two",
					ClassName:        "TableTwo",
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TablesFromConfig(tt.args.configFilepath)
			assert.Equal(t, tt.want, got)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestTable_Filename(t *testing.T) {
	tests := []struct {
		name string
		t    *Table
		want string
	}{
		{
			"uses FilenameOverride when present",
			&Table{
				TableName:        "table_ones",
				FilenameOverride: "table_ones_override",
			},
			"table_ones_override",
		},
		{
			"uses TableName when FilenameOverride is not available",
			&Table{
				TableName: "table_ones",
			},
			"table_ones",
		},
		{
			"uses TableName when FilenameOverride is blank",
			&Table{
				TableName:        "table_ones",
				FilenameOverride: "",
			},
			"table_ones",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.t.Filename()
			assert.Equal(t, tt.want, got)
		})
	}
}
