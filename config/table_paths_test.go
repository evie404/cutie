package config

import (
	"path/filepath"
	"testing"
)

func TestTable_SchemaTablesDirPath(t *testing.T) {
	type fields struct {
		TableName        string
		FilenameOverride string
		ModelClass       string
	}
	tests := []struct {
		name  string
		table Table
		want  string
	}{
		{
			"returns schema/tables by default",
			Table{},
			filepath.Join("schema", "tables"),
		},
		{
			"returns override if present",
			Table{
				schemaTablesDirOverride: filepath.Join("lol", "path"),
			},
			filepath.Join("lol", "path"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.table.SchemaTablesDirPath(); got != tt.want {
				t.Errorf("Table.SchemaTablesDirPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTable_SchemaQueriesDirPath(t *testing.T) {
	tests := []struct {
		name  string
		table Table
		want  string
	}{
		{
			"returns schema/queries by default",
			Table{},
			filepath.Join("schema", "queries"),
		},
		{
			"returns override if present",
			Table{
				schemaQueriesDirOverride: filepath.Join("lol", "path"),
			},
			filepath.Join("lol", "path"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.table.SchemaQueriesDirPath(); got != tt.want {
				t.Errorf("Table.SchemaQueriesDirPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTable_DbModelsDirPath(t *testing.T) {
	tests := []struct {
		name  string
		table Table
		want  string
	}{
		{
			"returns dbmodels/{{table name}} by default",
			Table{
				TableName: "hi",
			},
			filepath.Join("dbmodels", "hi"),
		},
		{
			"returns override if present",
			Table{
				TableName:           "hi",
				dbModelsDirOverride: filepath.Join("lol", "path"),
			},
			filepath.Join("lol", "path", "hi"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.table.DbModelsDirPath(); got != tt.want {
				t.Errorf("Table.DbModelsDirPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
