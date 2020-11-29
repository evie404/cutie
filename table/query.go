package table

import (
	"fmt"
	"os"
	"text/template"
)

const queryTemplateRaw = `-- name: Get{{.ModelClass}}ByID :one
SELECT * FROM {{.TableName}} WHERE id = $1 LIMIT 1;

`

func CreateQueryAll(tables []Table) error {
	var err error

	for _, table := range tables {
		err = CreateQuery(table)
		if err != nil {
			return fmt.Errorf("creating query file: %w", err)
		}
	}

	return nil
}

func CreateQuery(table Table) error {
	var err error
	_, err = os.Stat(table.QuerySQLPath())
	if err == nil {
		// don't create file if it alredy exists
		return nil

	} else if !os.IsNotExist(err) {
		return fmt.Errorf("statting file %s: %w", table.QuerySQLPath(), err)
	}

	err = os.MkdirAll(table.SchemaQueriesDirPath(), 0777)
	if err != nil {
		return fmt.Errorf("creating directory %s: %w", table.SchemaTablesDirPath(), err)
	}

	f, err := os.Create(table.QuerySQLPath())
	if err != nil {
		return fmt.Errorf("creating file %s: %w", table.QuerySQLPath(), err)
	}
	defer f.Close()

	queryTemplate, err := template.New("queryTemplate").Parse(queryTemplateRaw)
	if err != nil {
		return fmt.Errorf("creating parsing query template: %w", err)
	}

	err = queryTemplate.ExecuteTemplate(f, "queryTemplate", table)
	if err != nil {
		return fmt.Errorf("executing query template: %w", err)
	}

	return nil
}
