#!/bin/bash

# Database reset script for Treasury backend
# This script resets the PostgreSQL database and runs migrations

echo "🔄 Resetting Treasury database..."

# Check if Docker is running
if ! docker info > /dev/null 2>&1; then
    echo "Error: Docker is not running. Please start Docker and try again."
    exit 1
fi

# Check if postgres container exists and is running
if ! docker ps --format "table {{.Names}}" | grep -q "postgres"; then
    echo "❌ PostgreSQL container 'postgres' is not running."
    echo "Please start your docker-compose services first:"
    echo "  docker-compose up -d"
    exit 1
fi

echo "✅ Found running PostgreSQL container: postgres"

# Reset the database by dropping and recreating it
echo "Resetting database..."
# First, terminate all connections to the treasury database
docker exec postgres psql -U postgres -c "SELECT pg_terminate_backend(pid) FROM pg_stat_activity WHERE datname = 'treasury' AND pid <> pg_backend_pid();"
# Drop and recreate the database
docker exec postgres psql -U postgres -c "DROP DATABASE IF EXISTS treasury;"
docker exec postgres psql -U postgres -c "CREATE DATABASE treasury;"

if [ $? -eq 0 ]; then
    echo "✅ Database reset successfully!"
else
    echo "❌ Failed to reset database"
    exit 1
fi

# Run migrations
echo "Running database migrations..."
cd "$(dirname "$0")/.."
go run migrations/migrate.go

if [ $? -eq 0 ]; then
    echo "🎉 Database reset and migrations completed successfully!"
    echo ""
    echo "Database connection details:"
    echo "  Host: localhost"
    echo "  Port: 5432"
    echo "  Database: treasury"
    echo "  Username: postgres"
    echo "  Password: password"
    echo ""
    echo "Note: This script works with your docker-compose PostgreSQL container."
else
    echo "❌ Migration failed"
    exit 1
fi
