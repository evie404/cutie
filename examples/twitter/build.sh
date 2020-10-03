#!/bin/sh

psql $TWITTER_EXAMPLE_CONNECTION_STRING < schema.sql

(cd ../../ && go build cmd/main.go) && DATABASE_URL=$TWITTER_EXAMPLE_CONNECTION_STRING ../../main
