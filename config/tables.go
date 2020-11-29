package config

import (
	"fmt"

	"github.com/rickypai/cutie/table"
)

func TablesFromConfig(configFilepath string) ([]table.Table, error) {
	config, err := ParseConfigFromYAMLPath(configFilepath)
	if err != nil {
		return nil, fmt.Errorf("error parsing config file %s: %w", configFilepath, err)
	}

	return config.Tables, nil
}
