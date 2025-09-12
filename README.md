# Task List Application

A modern desktop task management application built with Go, PostgreSQL, and Vue.js using the Wails framework.

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
goose up
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