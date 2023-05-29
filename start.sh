#!/bin/sh

set -e

echo "run db migration"
/app/migrate -path /app/migration -database postgresql://root:secret@localhost:5432/todoapp?sslmode=disable -verbose up

echo "start the app"
exec "$@"