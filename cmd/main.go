package main

import (
	"log"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/rickypai/cutie/config"
	"github.com/rickypai/cutie/modelgen"
	"github.com/rickypai/cutie/table"
)

func main() {
	cfg, err := config.ParseConfigFromYAMLPath(".cutie.yaml")
	if err != nil {
		log.Fatalf("error loading config: %s", err)
	}

	pgConfig, err := pgx.ParseConfig(os.Getenv("POSTGRES_HOST") + "/" + cfg.Database)

	tables := cfg.Tables
	err = table.DumpAll(pgConfig, tables)
	if err != nil {
		log.Fatalf("error dumping tables: %s", err)
	}

	err = table.CreateQueryAll(tables)
	if err != nil {
		log.Fatalf("error create query files: %s", err)
	}

	err = modelgen.GenerateSQLCModels(tables)
	if err != nil {
		log.Fatalf("error generating SQLC models: %s", err)
	}

	err = modelgen.GenerateSQLCModelMocks(tables)
	if err != nil {
		log.Fatalf("error generating SQLC mocks: %s", err)
	}

	log.Println("All done!")
}
