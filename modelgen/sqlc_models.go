package modelgen

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/jinzhu/inflection"
	"github.com/rickypai/cutie/config"
)

func GenerateSQLCModels(tables []config.Table) error {
	// TODO: check sqlc version and warn

	err := GenerateSQLCConfig(tables)
	if err != nil {
		return fmt.Errorf("generating sqlc config: %w", err)
	}

	cmd := exec.Command(
		"sqlc",
		"generate",
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("generating sqlc models: %w", err)
	}

	err = renameSQLCModelNames(tables)
	if err != nil {
		return fmt.Errorf("generating renaming sqlc models: %w", err)
	}

	return nil
}

func renameSQLCModelNames(tables []config.Table) error {
	for _, table := range tables {
		if table.ModelClass == "" {
			continue
		}

		defaultName := sqlcDefaultModelClass(inflection.Singular(table.TableName))

		if table.ModelClass == defaultName {
			continue
		}

		cmd := exec.Command(
			"gorename",
			"-from",
			fmt.Sprintf("\"./%s\".%s", table.DbModelsDirPath(), defaultName),
			"-to",
			table.ModelClass,
			"--force",
		)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Run()
		if err != nil {
			return fmt.Errorf("renaming sqlc model for %s: %w", table.TableName, err)
		}
	}

	return nil
}

// from internal/codegen/golang/struct.go
func sqlcDefaultModelClass(name string) string {
	// if rename := settings.Rename[name]; rename != "" {
	// 	return rename
	// }
	out := ""
	for _, p := range strings.Split(name, "_") {
		if p == "id" {
			out += "ID"
		} else {
			out += strings.Title(p)
		}
	}

	return out
}

func GenerateSQLCModelMocks(tables []config.Table) error {
	for _, table := range tables {
		cmd := exec.Command(
			"mockgen",
			fmt.Sprintf("-source=%s", table.QuerierInterfaceGoPath()),
			fmt.Sprintf("-destination=%s", table.QuerierMockGoPath()),
		)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Run()
		if err != nil {
			return fmt.Errorf("generating sqlc mocks for table %s: %w", table.TableName, err)
		}
	}

	return nil
}
