setup_db:
	psql -Atx ${POSTGRES_HOST}/testdb cutie_test < schema.sql

build:
	rm -rf dbmodels
	rm -rf schema/tables
	(cd ../../ && go build cmd/main.go)
	(../../main && go build ./...)
