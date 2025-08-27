#!/bin/bash

# Database setup script for Treasury backend
# This script sets up a PostgreSQL database using Docker

echo "Setting up PostgreSQL database for Treasury backend..."

# Check if Docker is running
if ! docker info > /dev/null 2>&1; then
    echo "Error: Docker is not running. Please start Docker and try again."
    exit 1
fi

# Check if postgres-treasury container already exists
if docker ps -a --format "table {{.Names}}" | grep -q "postgres-treasury"; then
    echo "PostgreSQL container 'postgres-treasury' already exists."
    echo "Removing existing container..."
    docker rm -f postgres-treasury
fi

# Run PostgreSQL container
echo "Starting PostgreSQL container..."
docker run --name postgres-treasury \
  -e POSTGRES_PASSWORD=password \
  -e POSTGRES_DB=treasury \
  -p 5432:5432 \
  -d postgres:15

# Wait for database to be ready
echo "Waiting for database to be ready..."
sleep 5

# Check if container is running
if docker ps --format "table {{.Names}}" | grep -q "postgres-treasury"; then
    echo "✅ PostgreSQL database is running!"
    echo "Database connection details:"
    echo "  Host: localhost"
    echo "  Port: 5432"
    echo "  Database: treasury"
    echo "  Username: postgres"
    echo "  Password: password"
    echo ""
    echo "You can now run the Go backend application:"
    echo "  cd apps/web-backend"
    echo "  go run main.go"
else
    echo "❌ Failed to start PostgreSQL container"
    exit 1
fi
