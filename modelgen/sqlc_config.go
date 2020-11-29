package modelgen

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/rickypai/cutie/config"
)

type v1GenerateSettings struct {
	Version  string              `json:"version" yaml:"version"`
	Packages []v1PackageSettings `json:"packages" yaml:"packages"`
	Rename   map[string]string   `json:"rename,omitempty" yaml:"rename,omitempty"`
}

type v1PackageSettings struct {
	Name          string `json:"name" yaml:"name"`
	Engine        string `json:"engine,omitempty" yaml:"engine"`
	Path          string `json:"path" yaml:"path"`
	Schema        string `json:"schema" yaml:"schema"`
	Queries       string `json:"queries" yaml:"queries"`
	EmitInterface bool   `json:"emit_interface" yaml:"emit_interface"`
}

func GenerateSQLCConfig(tables []config.Table) error {
	var err error

	err = generateSQLCModelsDir(tables)
	if err != nil {
		return fmt.Errorf("generating directories: %w", err)
	}

	config := generateSQLCConfig(tables)

	configJson, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return fmt.Errorf("marshalling json: %w", err)
	}

	err = ioutil.WriteFile("sqlc.json", configJson, 0644)
	if err != nil {
		return fmt.Errorf("writing json file: %w", err)
	}

	return nil
}

func generateSQLCModelsDir(tables []config.Table) error {
	var err error

	for _, table := range tables {
		err = os.MkdirAll(table.DbModelsDirPath(), 0777)
		if err != nil {
			return fmt.Errorf("creating directory %s: %w", table.DbModelsDirPath(), err)
		}
	}

	return nil
}

func generateSQLCConfig(tables []config.Table) v1GenerateSettings {
	config := v1GenerateSettings{
		Version: "1",
	}

	for _, table := range tables {
		config.Packages = append(config.Packages, v1PackageSettings{
			Schema:        table.SchemaSQLPath(),
			Queries:       table.QuerySQLPath(),
			Engine:        "postgresql",
			EmitInterface: true,
			Name:          table.Filename(),
			Path:          table.DbModelsDirPath(),
		})
	}

	return config
}
