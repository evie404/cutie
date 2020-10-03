package table

import (
	"fmt"
	"io/ioutil"

	yaml "github.com/goccy/go-yaml"
)

func TablesFromConfig(configFilepath string) ([]Table, error) {
	var tables []Table

	data, err := ioutil.ReadFile(configFilepath)
	if err != nil {
		return nil, fmt.Errorf("reading file: %w", err)
	}

	err = yaml.Unmarshal(data, &tables)
	if err != nil {
		return nil, fmt.Errorf("unmarshalling yaml: %w", err)
	}

	return tables, nil
}

type Table struct {
	TableName        string `yaml:"table_name"`
	FilenameOverride string `yaml:"filename_override"`
	ClassName        string `yaml:"class_name"`
}

func (t *Table) Filename() string {
	if t.FilenameOverride != "" {
		return t.FilenameOverride
	}

	return t.TableName
}

func (t *Table) SQLFilename() string {
	return fmt.Sprintf("%s.sql", t.Filename())
}
