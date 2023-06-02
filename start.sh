#!/bin/sh

set -e

echo "run db migration"
source /app/git checkoutapp.env
/app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up

echo "start the app"
exec "$@"