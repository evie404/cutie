package config

import (
	"path/filepath"
)

func (t *Table) SchemaSQLPath() string {
	return filepath.Join(t.SchemaTablesDirPath(), t.SQLFilename())
}

func (t *Table) QuerySQLPath() string {
	return filepath.Join(t.SchemaQueriesDirPath(), t.SQLFilename())
}

func (t *Table) SchemaTablesDirPath() string {
	if t.schemaTablesDirOverride != "" {
		return t.schemaTablesDirOverride
	}

	return filepath.Join("schema", "tables")
}

func (t *Table) SchemaQueriesDirPath() string {
	if t.schemaQueriesDirOverride != "" {
		return t.schemaQueriesDirOverride
	}

	return filepath.Join("schema", "queries")
}

func (t *Table) DbModelsDirPath() string {
	if t.dbModelsDirOverride != "" {
		return t.dbModelsDirOverride
	}

	return filepath.Join("dbmodels", t.Filename())
}

func (t *Table) QuerierInterfaceGoPath() string {
	return filepath.Join(t.DbModelsDirPath(), "querier.go")
}

func (t *Table) QuerierMockGoPath() string {
	return filepath.Join(t.DbModelsDirPath(), "mock_"+t.Filename(), "querier_mock.go")
}

func (t *Table) ModelsDirPath() string {
	return filepath.Join("models", t.Filename())
}
