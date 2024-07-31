#!/bin/bash
set -e

# Check if the user already exists
if psql -U "$POSTGRES_USER" -d "postgres" -tAc "SELECT 1 FROM pg_roles WHERE rolname='$POSTGRES_USER'" | grep -q 1; then
  echo "User 'robin_admin' already exists, skipping user creation."
else
  psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "postgres" <<EOSQL
      CREATE USER $POSTGRES_USER WITH PASSWORD '$POSTGRES_PASSWORD';
EOSQL
fi

# Check if the database already exists
if psql -U "$POSTGRES_USER" -d "postgres" -tAc "SELECT 1 FROM pg_database WHERE datname='$POSTGRES_DB'" | grep -q 1; then
  echo "Database '$POSTGRES_DB' already exists, skipping database creation."
else
  psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "postgres" <<EOSQL
      CREATE DATABASE $POSTGRES_DB;
      GRANT ALL PRIVILEGES ON DATABASE $POSTGRES_DB TO $POSTGRES_USER;
EOSQL
fi
