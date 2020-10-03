package table

import (
	"fmt"
	"io/ioutil"

	yaml "github.com/goccy/go-yaml"
	"go.uber.org/multierr"
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

	for _, table := range tables {
		if table.IsValid() {
			continue
		}

		return nil, fmt.Errorf("invalid config for table `%s`: %w", table.TableName, table.InvalidReasons())
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

func (t *Table) IsValid() bool {
	return t.InvalidReasons() == nil
}

func (t *Table) InvalidReasons() error {
	var errs []error

	if t.TableName == "" {
		errs = append(errs, fmt.Errorf("table_name cannot be blank"))
	}

	if t.FilenameOverride != "" && t.TableName == t.FilenameOverride {
		errs = append(errs, fmt.Errorf("filename_override `%s` cannot be same as table_name", t.FilenameOverride))
	}

	return multierr.Combine(errs...)
}
