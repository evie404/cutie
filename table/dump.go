package table

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/jackc/pgconn"
)

func DumpAll(dbConfig *pgconn.Config, tables []Table) error {
	var err error

	for _, table := range tables {
		err = DumpOne(dbConfig, table)
		if err != nil {
			return fmt.Errorf("dumping tables: %w", err)
		}
	}

	return nil
}

func DumpOne(dbConfig *pgconn.Config, table Table) error {
	var err error

	err = os.MkdirAll(table.SchemaTablesDirPath(), 0644)
	if err != nil {
		return fmt.Errorf("creating directory %s: %w", table.SchemaTablesDirPath(), err)
	}

	cmd := exec.Command(
		"pg_dump",
		"--schema-only",
		"--no-comments",
		fmt.Sprintf("--host=%s", dbConfig.Host),
		fmt.Sprintf("--username=%s", dbConfig.User),
		fmt.Sprintf("--dbname=%s", dbConfig.Database),
		fmt.Sprintf("--table=%s", table.TableName),
		fmt.Sprintf("--file=%s", table.SchemaSQLPath()),
	)
	cmd.Env = append(cmd.Env, fmt.Sprintf("PGPASSWORD=%s", dbConfig.Password))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("dumping table %s: %w", table.TableName, err)
	}

	return nil
}
