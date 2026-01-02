# Go Starter Kit

[![Go Version](https://img.shields.io/badge/Go-1.25.5-blue.svg)](https://golang.org)
[![codecov](https://codecov.io/gh/cristiano-pacheco/go-starter-kit/graph/badge.svg?token=NPIL8D7H6D)](https://codecov.io/gh/cristiano-pacheco/go-starter-kit)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)

A scalable Go API starter kit with clean architecture principles, authentication, event-driven patterns, and observability.

## Features

- 🔐 JWT Authentication & Authorization
- 📧 Email verification & templates
- 🔄 Event-driven with Apache Kafka
- 💾 Redis caching
- 📊 OpenTelemetry, Jaeger & Prometheus
- 📝 Swagger documentation
- 🐳 Docker Compose ready
- 🧪 Unit tests
- 🗄️ Database migrations
- 🎯 Dependency injection (Uber FX)
- 🌐 Fiber HTTP framework

## Quick Start

```bash
# Clone
git clone https://github.com/cristiano-pacheco/go-starter-kit.git
cd go-starter-kit

# Start infrastructure
docker-compose up -d

# Run migrations
make migrate

# Start server
make run
```

Server runs at `http://localhost:9000`  
Swagger docs at `http://localhost:9000/swagger/index.html`

## Requirements

- Go 1.25.5+
- Docker & Docker Compose
- Make

## Configuration

Create a `.env` file:

```bash
# Database
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=go-starter-kit

# Redis
REDIS_HOST=localhost:6379

# Kafka
KAFKA_BROKERS=localhost:9092

# JWT
JWT_SECRET=your-secret-key
JWT_EXPIRATION=24h

# Server
SERVER_PORT=9000
```

## Project Structure

```
├── cmd/                    # CLI commands
├── internal/
│   ├── modules/           # Feature modules (identity, etc)
│   └── shared/            # Shared code
├── pkg/                   # Reusable packages
├── migrations/            # SQL migrations
└── docs/                  # Swagger docs
```

## Commands

```bash
# Development
make run                   # Run server
make migrate              # Run migrations

# Testing
make test                 # Run unit tests
make cover               # Coverage report

# Code quality
make lint                # Lint code
make static              # Run all checks
make update-swagger      # Update docs
```

## Tech Stack

**Core:** Go, Fiber, GORM, Viper, Cobra  
**Database:** PostgreSQL, Redis, golang-migrate  
**Messaging:** Apache Kafka  
**Observability:** OpenTelemetry, Jaeger, Prometheus, zerolog  
**Security:** JWT, bcrypt  
**Testing:** testify, mockery  
**DI:** Uber FX

## Docker Services

Docker Compose starts:
- PostgreSQL (5432)
- Redis (6379)
- Kafka (9092)
- Kafka UI (8080)
- Jaeger (16686)
- Prometheus (9090)

## License

MIT License - see [LICENSE](LICENSE)

## Author

**Cristiano Pacheco** - [@cristiano-pacheco](https://github.com/cristiano-pacheco)
