#!/bin/bash
source dev.env

sleep 2 && goose -dir "${MIGRATION_DIR}" postgres "host=${PG_HOST} user=${PG_USER} password=${PG_PASSWORD} dbname=${PG_DATABASE_NAME} port=${PG_PORT_OUTER}" up -v