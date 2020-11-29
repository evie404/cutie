setup_db:
	psql -Atx ${QT_DATABASE_URL} cutie_test < schema.sql

build:
	rm -rf dbmodels
	rm -rf schema/tables
	(cd ../../ && go build cmd/main.go)
	(DATABASE_URL=${QT_DATABASE_URL} ../../main && go build ./...)
