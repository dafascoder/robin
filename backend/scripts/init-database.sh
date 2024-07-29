#!/bin/bash
set -e

# Check if the user already exists
if psql -U "$POSTGRES_USER" -d "postgres" -tAc "SELECT 1 FROM pg_roles WHERE rolname='voltig'" | grep -q 1; then
  echo "User 'voltig' already exists, skipping user creation."
else
  psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "postgres" <<EOSQL
      CREATE USER voltig WITH PASSWORD 'dummy';
EOSQL
fi

# Check if the database already exists
if psql -U "$POSTGRES_USER" -d "postgres" -tAc "SELECT 1 FROM pg_database WHERE datname='voltigDB'" | grep -q 1; then
  echo "Database 'voltigDB' already exists, skipping database creation."
else
  psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "postgres" <<EOSQL
      CREATE DATABASE voltigDB;
      GRANT ALL PRIVILEGES ON DATABASE voltigDB TO voltig;
EOSQL
fi
