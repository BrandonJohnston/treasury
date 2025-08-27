# Treasury

A full-stack web application with a Next.js frontend and Go backend with PostgreSQL database.

## Requirements

### Frontend
- Node.js (minimum v22.18.0)
  - https://nodejs.org/en/download
- nvm (Node Version Manager)
  - https://github.com/nvm-sh/nvm?tab=readme-ov-file#installing-and-updating

### Backend
- Go 1.25 or higher
  - https://golang.org/dl/
- PostgreSQL database
- Docker (optional, for containerized development)

## Quick Start

### 1. Database Setup

#### Option A: Using Docker (Recommended for development)
```bash
cd apps/web-backend
./scripts/setup-db.sh
```

#### Option B: Local PostgreSQL
1. Install PostgreSQL on your system
2. Create a database:
   ```sql
   CREATE DATABASE treasury;
   ```

### 2. Backend Setup

```bash
cd apps/web-backend

# Install dependencies
go mod tidy

# Create .env file (optional, uses defaults if not present)
echo "DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=treasury
DB_SSLMODE=disable" > .env

# Run the backend
go run main.go
```

The backend will be available at `http://localhost:8080`

### 3. Frontend Setup

```bash
cd apps/web-frontend

# Install dependencies
npm install

# Run the development server
npm run dev
```

The frontend will be available at `http://localhost:3000`

## Project Structure

```
treasury/
├── apps/
│   ├── web-backend/          # Go backend with PostgreSQL
│   │   ├── config/           # Configuration management
│   │   ├── database/         # Database connection and setup
│   │   ├── handlers/         # HTTP request handlers
│   │   ├── models/           # Data models
│   │   ├── repository/       # Database operations
│   │   ├── routes/           # Route definitions
│   │   ├── scripts/          # Setup scripts
│   │   └── main.go           # Application entry point
│   └── web-frontend/         # Next.js frontend
│       ├── src/
│       │   ├── app/          # Next.js app router
│       │   ├── components/   # React components
│       │   └── lib/          # Utility functions
│       └── package.json
├── docker-compose.yml        # Docker services configuration
└── README.md
```

## API Endpoints

### Authentication
- `GET /api/auth/user` - Get all users
- `POST /api/auth/login` - User login
- `POST /api/auth/register` - User registration

### Example API Usage

```bash
# Register a new user
curl -X POST http://localhost:8080/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "password123",
    "name": "John Doe"
  }'

# Login
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "password123"
  }'

# Get all users
curl http://localhost:8080/api/auth/user
```

## Development

### Running with Docker Compose

```bash
# Start all services
docker-compose up -d

# View logs
docker-compose logs -f

# Stop all services
docker-compose down
```

### Database Management

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
