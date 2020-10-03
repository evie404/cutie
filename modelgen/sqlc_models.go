package modelgen

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/rickypai/cutie/table"
)

func GenerateSQLCModels(tables []table.Table) error {
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

	return nil
}

func GenerateSQLCModelMocks(tables []table.Table) error {
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
			return fmt.Errorf("generating sqlc mocks: %w", err)
		}
	}

	return nil
}
