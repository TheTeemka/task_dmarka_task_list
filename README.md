# Task List Application

A modern desktop task management application built with Go, PostgreSQL, and Vue.js using the Wails framework.

## Features

- ✅ Task creation and management
- 🏷️ Tag-based organization
- 📊 Priority levels
- 📅 Due dates and completion tracking
- 🔍 Advanced filtering and search
- 🎨 Modern, responsive UI with Tailwind CSS
- 🐳 Docker support for easy deployment
- 📱 Cross-platform desktop application

## Prerequisites

Before you begin, ensure you have the following installed:

- **Go** (version 1.19 or later)
- **Node.js** (version 16 or later) and npm
- **Docker** and Docker Compose (for containerized setup)
- **PostgreSQL** (if running without Docker)

## Quick Start with Docker

The fastest way to get started is using Docker Compose:

### 1. Clone the repository
```bash
git clone <your-repo-url>
cd task_dmarka_task_list
```

### 2. Start PostgreSQL database
```bash
docker-compose up -d postgres
```

### 3. Run database migrations
```bash
# Install Goose (if not already installed)
go install github.com/pressly/goose/v3/cmd/goose@latest

# Run migrations
goose -dir ./internal/database/migrations postgres "postgres://postgres:postgres@localhost:5432/tasklist?sslmode=disable" up
```

### 4. Install frontend dependencies
```bash
npm install
```

### 5. Build and run the application
```bash
# For development
wails dev

# For production build
wails build
```

## Manual Setup (Without Docker)

If you prefer to set up PostgreSQL manually:

### 1. Install and configure PostgreSQL
```bash
# Install PostgreSQL (macOS with Homebrew)
brew install postgresql
brew services start postgresql

# Create database
createdb tasklist
```

### 2. Configure environment variables
Copy `.env` and update the database connection:
```bash
cp .env .env.local
# Edit .env.local with your PostgreSQL credentials
```

### 3. Run migrations
```bash
goose -dir ./internal/database/migrations postgres "your-connection-string" up
```

### 4. Install dependencies and run
```bash
# Backend dependencies
go mod download

# Frontend dependencies
npm install

# Run in development mode
wails dev
```

## Development

### Project Structure
```
├── main.go                 # Application entry point
├── app.go                  # Wails application setup
├── wails.json             # Wails configuration
├── docker-compose.yml     # Docker services
├── .env                   # Environment variables
├── init.sql              # Database initialization
├── frontend/             # Vue.js frontend
│   ├── src/
│   ├── package.json
│   └── vite.config.js
├── internal/
│   ├── config/           # Configuration management
│   ├── database/         # Database connection & migrations
│   ├── models/           # Data models
│   ├── repository/       # Data access layer
│   └── service/          # Business logic
└── build/                # Build artifacts
```

### Available Commands

```bash
# Development
wails dev              # Start development server with hot reload
npm run dev           # Start frontend dev server only

# Building
wails build           # Build production application
wails build -clean    # Clean build artifacts

# Database
goose up              # Run database migrations
goose down            # Rollback migrations
goose status          # Check migration status

# Docker
docker-compose up -d          # Start all services
docker-compose down           # Stop all services
docker-compose logs postgres  # View PostgreSQL logs
```

### Environment Variables

Create a `.env` file with the following variables:

```env
# Database Configuration
PSQL_HOST=localhost
PSQL_PORT=5432
PSQL_USER=postgres
PSQL_PASSWORD=your_password
PSQL_NAME=tasklist
PSQL_SSLMODE=disable
PSQL_SOURCE=postgres://postgres:your_password@localhost:5432/tasklist?sslmode=disable

# Goose Migration Configuration
GOOSE_DBSTRING=${PSQL_SOURCE}
GOOSE_DRIVER=postgres
GOOSE_MIGRATION_DIR=./internal/database/migrations
```

## Building for Production

### Desktop Application
```bash
# Build for current platform
wails build

# Build for specific platforms
wails build -platform windows/amd64
wails build -platform darwin/universal
wails build -platform linux/amd64
```

### Docker Deployment
```bash
# Build and run with Docker Compose
docker-compose up --build

# Or build image manually
docker build -t tasklist-app .
docker run -p 8080:8080 tasklist-app
```

## Database Schema

The application uses PostgreSQL with the following main tables:

- **tasks**: Main task records with title, description, status, priority, dates
- **tags**: Tag definitions with name and color
- **task_tags**: Many-to-many relationship between tasks and tags

### Running Migrations

```bash
# Check current migration status
goose -dir ./internal/database/migrations postgres "${PSQL_SOURCE}" status

# Apply new migrations
goose -dir ./internal/database/migrations postgres "${PSQL_SOURCE}" up

# Rollback last migration
goose -dir ./internal/database/migrations postgres "${PSQL_SOURCE}" down
```

## Troubleshooting

### Common Issues

1. **Database connection failed**
   - Ensure PostgreSQL is running
   - Check connection string in `.env`
   - Verify database exists: `createdb tasklist`

2. **Build fails**
   - Run `go mod tidy` to update dependencies
   - Clear build cache: `wails build -clean`
   - Check Go version: `go version`

3. **Frontend not loading**
   - Run `npm install` in frontend directory
   - Clear node modules: `rm -rf node_modules && npm install`

4. **Docker issues**
   - Check Docker Desktop is running
   - Clear Docker cache: `docker system prune`
   - Rebuild containers: `docker-compose up --build`

### Logs and Debugging

```bash
# View application logs
wails dev  # Shows both frontend and backend logs

# View database logs
docker-compose logs postgres

# Debug database connection
go run -tags debug main.go
```

## Contributing

1. Fork the repository
2. Create a feature branch: `git checkout -b feature/your-feature`
3. Make your changes and test thoroughly
4. Run migrations if schema changes: `goose up`
5. Commit your changes: `git commit -am 'Add new feature'`
6. Push to the branch: `git push origin feature/your-feature`
7. Submit a pull request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Support

For questions or issues:
- Create an issue on GitHub
- Check the Wails documentation: https://wails.io/docs
- PostgreSQL documentation: https://www.postgresql.org/docs/

---

Built with ❤️ using [Wails](https://wails.io), [Vue.js](https://vuejs.org), and [PostgreSQL](https://postgresql.org)
