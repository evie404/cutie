package config

import (
	"fmt"
	"math/rand"
	"path/filepath"
	"testing"

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
			"errors on file with no database",
			args{
				filepath.Join("testdata", "no_database.yaml"),
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
				Database: "lol",
				Tables: []Table{
					{
						TableName:        "table_ones",
						FilenameOverride: "table_one",
						ModelClass:       "TableOne",
					},
					{
						TableName:        "table_twos",
						FilenameOverride: "table_two",
						ModelClass:       "TableTwo",
					},
				},
			},
			false,
		},
		{
			"load tables with dir override",
			args{
				filepath.Join("testdata", "tables_with_dir_override.yaml"),
			},
			&Config{
				Database: "lol",
				Tables: []Table{
					{
						TableName:                "table_ones",
						FilenameOverride:         "table_one",
						ModelClass:               "TableOne",
						schemaTablesDirOverride:  "lol/schema_tables_dir",
						schemaQueriesDirOverride: "lol/schema_queries_dir",
						dbModelsDirOverride:      "lol/dbmodels_dir",
					},
					{
						TableName:                "table_twos",
						FilenameOverride:         "table_two",
						ModelClass:               "TableTwo",
						schemaTablesDirOverride:  "lol/schema_tables_dir",
						schemaQueriesDirOverride: "lol/schema_queries_dir",
						dbModelsDirOverride:      "lol/dbmodels_dir",
					},
				},
				SchemaTablesDirOverride:  "lol/schema_tables_dir",
				SchemaQueriesDirOverride: "lol/schema_queries_dir",
				DbModelsDirOverride:      "lol/dbmodels_dir",
			},
			false,
		},
		{
			"load tables_with_enums.yaml",
			args{
				filepath.Join("testdata", "tables_with_enums.yaml"),
			},
			&Config{
				Database: "lol",
				Tables: []Table{
					{
						TableName:        "table_ones",
						FilenameOverride: "table_one",
						ModelClass:       "TableOne",
						Enums: []Enum{
							{
								EnumName: "lol",
							},
							{
								EnumName: "hi",
							},
						},
					},
					{
						TableName:        "table_twos",
						FilenameOverride: "table_two",
						ModelClass:       "TableTwo",
						Enums: []Enum{
							{
								EnumName: "twos",
							},
						},
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
