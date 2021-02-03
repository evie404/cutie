package table

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/jackc/pgx/v4"
	"github.com/rickypai/cutie/config"
)

func DumpAll(dbConfig *pgx.ConnConfig, tables []config.Table) error {
	var err error

	for _, table := range tables {
		err = DumpOne(dbConfig, table)
		if err != nil {
			return fmt.Errorf("dumping tables: %w", err)
		}
	}

	return nil
}

func DumpOne(dbConfig *pgx.ConnConfig, table config.Table) error {
	return dumpSchema(dbConfig, table)
}

func dumpSchema(dbConfig *pgx.ConnConfig, table config.Table) error {
	var err error

	err = os.MkdirAll(table.SchemaTablesDirPath(), 0777)
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

	dumpEnums(dbConfig, table)

	// TODO: remove comments

	return nil
}

func dumpEnums(dbConfig *pgx.ConnConfig, table config.Table) error {
	ctx := context.Background()
	conn, err := pgx.ConnectConfig(ctx, dbConfig)
	if err != nil {
		return fmt.Errorf("connecting to table %s: %w", table.TableName, err)
	}
	defer conn.Close(ctx)

	for _, enum := range table.Enums {
		var values []string

		query := fmt.Sprintf(
			`SELECT unnest(enum_range(NULL::%s))::text AS values;`, // TODO: escape this
			enum.EnumName,
		)

		rows, err := conn.Query(ctx, query)
		if err != nil {
			return fmt.Errorf("querying enum %s with query `%s`: %w", enum.EnumName, query, err)
		}

		for rows.Next() {
			var value string
			err = rows.Scan(&value)
			if err != nil {
				return fmt.Errorf("querying enum %s value row: %w", enum.EnumName, query, err)
			}

			values = append(values, fmt.Sprintf("'%s'", value))
		}
		err = rows.Err()
		if err != nil {
			return fmt.Errorf("querying enum %s value row: %w", enum.EnumName, query, err)
		}

		createStatement := fmt.Sprintf(`CREATE TYPE %s AS ENUM (%s);`, enum.EnumName, strings.Join(values, ", "))

		f, err := os.OpenFile(table.SchemaSQLPath(), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return fmt.Errorf("opening file %s: %w", table.SchemaSQLPath(), err)
		}
		defer f.Close()

		_, err = f.WriteString(fmt.Sprintf("\n%s\n", createStatement))

		if err != nil {
			return fmt.Errorf("writing text to file %s: %w", table.SchemaSQLPath(), err)
		}
	}

	return nil
}
