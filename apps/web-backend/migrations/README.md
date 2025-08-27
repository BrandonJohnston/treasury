# Database Migrations

This directory contains database migrations for the Treasury application.

## Migration Files

Migrations are SQL files named with a numeric prefix (e.g., `001_create_users_table.sql`). The numeric prefix determines the order in which migrations are applied.

## Available Commands

### From the root directory:

```bash
# Reset the database (drops and recreates the treasury database, then runs all migrations)
npm run db:reset

# Run migrations only (requires database to be running)
npm run db:migrate


```

### From the backend directory:

```bash
# Run migrations
go run migrations/migrate.go

# Reset database (requires docker-compose to be running)
./scripts/reset-db.sh


```

## Migration System Features

- **Automatic ordering**: Migrations are applied in order based on their numeric prefix
- **Idempotent**: Running migrations multiple times is safe
- **Transaction safety**: Each migration runs in a transaction
- **Tracking**: Applied migrations are tracked in a `migrations` table
- **Reset functionality**: Drops and recreates the treasury database with fresh data

## Creating New Migrations

1. Create a new SQL file in the `migrations` directory
2. Name it with the next sequential number (e.g., `002_add_user_roles.sql`)
3. Write your SQL statements
4. Run `npm run db:migrate` to apply the migration

## Example Migration

```sql
-- Migration: 002_add_user_roles.sql
-- Description: Adds role column to users table
-- Created: 2024-01-02

ALTER TABLE users ADD COLUMN role VARCHAR(50) DEFAULT 'user';
CREATE INDEX idx_users_role ON users(role);
```

## Database Connection Details

- **Host**: localhost
- **Port**: 5432
- **Database**: treasury
- **Username**: postgres
- **Password**: password

## Troubleshooting

If you encounter issues:

1. Make sure Docker is running
2. Check that the PostgreSQL container is running: `docker ps`
3. Verify database connection: `docker exec -it postgres-treasury psql -U postgres -d treasury`
4. Check migration status: `SELECT * FROM migrations ORDER BY id;`
