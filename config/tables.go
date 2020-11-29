package config

import (
	"fmt"
	"io/ioutil"

	"github.com/goccy/go-yaml"
	"github.com/rickypai/cutie/table"
)

func TablesFromConfig(configFilepath string) ([]table.Table, error) {
	var tables []table.Table

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

	if len(tables) == 0 {
		return nil, fmt.Errorf("no tables found in %s", configFilepath)
	}

	return tables, nil
}
