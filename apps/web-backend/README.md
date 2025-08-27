# Web Backend

A Go backend application built with Fiber framework and PostgreSQL database.

## Features

- User authentication (login/register)
- PostgreSQL database integration
- RESTful API endpoints
- CORS support for frontend integration

## Prerequisites

- Go 1.25 or higher
- PostgreSQL database
- Docker (optional, for containerized development)

## Database Setup

### Option 1: Local PostgreSQL

1. Install PostgreSQL on your system
2. Create a database named `treasury`:
   ```sql
   CREATE DATABASE treasury;
   ```

### Option 2: Docker PostgreSQL

Run PostgreSQL using Docker:
```bash
docker run --name postgres-treasury \
  -e POSTGRES_PASSWORD=password \
  -e POSTGRES_DB=treasury \
  -p 5432:5432 \
  -d postgres:15
```

## Environment Configuration

Create a `.env` file in the `apps/web-backend` directory with the following variables:

```env
# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=treasury
DB_SSLMODE=disable

# Server Configuration
PORT=8080
```

## Installation

1. Navigate to the backend directory:
   ```bash
   cd apps/web-backend
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Run the application:
   ```bash
   go run main.go
   ```

The server will start on `http://localhost:8080`

## API Endpoints

### Authentication

- `GET /api/auth/user` - Get all users
- `POST /api/auth/login` - User login
- `POST /api/auth/register` - User registration

### Example Usage

#### Register a new user:
```bash
curl -X POST http://localhost:8080/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "password123",
    "name": "John Doe"
  }'
```

#### Login:
```bash
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "password123"
  }'
```

#### Get all users:
```bash
curl http://localhost:8080/api/auth/user
```

## Project Structure

```
apps/web-backend/
├── config/          # Configuration management
├── database/        # Database connection and setup
├── handlers/        # HTTP request handlers
├── models/          # Data models
├── repository/      # Database operations
├── routes/          # Route definitions
├── main.go          # Application entry point
└── go.mod           # Go module file
```

## Development

### Running with Docker

Use the provided Dockerfile for containerized development:

```bash
docker build -f Dockerfile.dev -t treasury-backend .
docker run -p 8080:8080 treasury-backend
```

### Database Migrations

The application automatically creates the necessary database tables on startup. The `users` table includes:

- `id` (SERIAL PRIMARY KEY)
- `email` (VARCHAR(255) UNIQUE)
- `password_hash` (VARCHAR(255))
- `name` (VARCHAR(255))
- `created_at` (TIMESTAMP)
- `updated_at` (TIMESTAMP)

## Security Notes

⚠️ **Important**: This is a development setup. For production:

1. Use strong password hashing (bcrypt, Argon2)
2. Implement JWT or session-based authentication
3. Use HTTPS
4. Set up proper database security
5. Use environment-specific configuration
6. Implement rate limiting and input validation
