#!/bin/bash

# Function to check if the database is ready
wait_for_db() {
    until psql "postgres://postgres:qwerty@db:5432/postgres?sslmode=disable" -c '\q'; do
        >&2 echo "Waiting for database to be ready..."
        sleep 5
    done
    >&2 echo "Database is ready!"
}

# Run the function to wait for the database
wait_for_db

# Run migrations
echo "Applying database migrations..."
migrate -path ./schema -database 'postgres://postgres:qwerty@db:5432/postgres?sslmode=disable' up

# Exit with the status of the migration command
exit $?
