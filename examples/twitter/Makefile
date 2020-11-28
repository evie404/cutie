setup_db:
	psql -Atx ${QT_DATABASE_URL} cutie_test < schema.sql

build:
	(cd ../../ && go build cmd/main.go)
	(DATABASE_URL=${QT_DATABASE_URL} ../../main && go build ./...)
