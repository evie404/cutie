package config

import (
	"fmt"
)

func TablesFromConfig(configFilepath string) ([]Table, error) {
	config, err := ParseConfigFromYAMLPath(configFilepath)
	if err != nil {
		return nil, fmt.Errorf("error parsing config file %s: %w", configFilepath, err)
	}

	return config.Tables, nil
}
