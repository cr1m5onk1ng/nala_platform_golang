#!/bin/sh

set -e

echo "run db migration"
source /nala/app.env
/nala/migrate -path /nala/migration -database "$DB_SOURCE" -verbose up

echo "start the app"
exec "$@"