package config

import (
	"fmt"
	"math/rand"
	"path/filepath"
	"testing"

	"github.com/rickypai/cutie/table"
	"github.com/stretchr/testify/assert"
)

func TestParseConfigFromYAMLPath(t *testing.T) {
	type args struct {
		configFilepath string
	}
	tests := []struct {
		name    string
		args    args
		want    *Config
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
			"errors on file with no tables",
			args{
				filepath.Join("testdata", "no_tables.yaml"),
			},
			nil,
			true,
		},
		{
			"load tables1.yaml",
			args{
				filepath.Join("testdata", "tables1.yaml"),
			},
			&Config{
				Tables: []table.Table{
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
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseConfigFromYAMLPath(tt.args.configFilepath)
			assert.Equal(t, tt.want, got)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
