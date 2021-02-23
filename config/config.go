package config

import (
	"fmt"
	"io/ioutil"

	"github.com/goccy/go-yaml"
)

type Config struct {
	Database string `yaml:"database"`

	Tables []Table `yaml:"tables"`

	SchemaTablesDirOverride  string `yaml:"schema_tables_dir"`
	SchemaQueriesDirOverride string `yaml:"schema_queries_dir"`
	DbModelsDirOverride      string `yaml:"dbmodels_dir"`
}

func ParseConfigFromYAMLPath(configFilepath string) (*Config, error) {
	var config Config

	data, err := ioutil.ReadFile(configFilepath)
	if err != nil {
		return nil, fmt.Errorf("reading file: %w", err)
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("unmarshalling yaml: %w", err)
	}

	for i, table := range config.Tables {
		if table.IsValid() {
			config.Tables[i].schemaTablesDirOverride = config.SchemaTablesDirOverride
			config.Tables[i].schemaQueriesDirOverride = config.SchemaQueriesDirOverride
			config.Tables[i].dbModelsDirOverride = config.DbModelsDirOverride

			continue
		}

		return nil, fmt.Errorf("invalid config for table `%s`: %w", table.TableName, table.InvalidReasons())
	}

	if len(config.Tables) == 0 {
		return nil, fmt.Errorf("no tables found in %s", configFilepath)
	}

	return &config, nil
}
