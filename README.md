# EventHub — Backend API

![Go](https://img.shields.io/badge/Go-1.27.1-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Gin](https://img.shields.io/badge/Gin-1.12.0-008ECF?style=for-the-badge)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-16-4169E1?style=for-the-badge&logo=postgresql&logoColor=white)
![Redis](https://img.shields.io/badge/Redis-7-DC382D?style=for-the-badge&logo=redis&logoColor=white)
![Swagger](https://img.shields.io/badge/Swagger-Swaggo-85EA2D?style=for-the-badge&logo=swagger&logoColor=black)
![Docker](https://img.shields.io/badge/Docker-Multi--stage-2496ED?style=for-the-badge&logo=docker&logoColor=white)

**EventHub Backend** is a RESTful API built with Golang and Gin for an event and community management platform.

It handles authentication, event management, community membership, user profiles, saved events, discussions, and related application data.

Developed as part of the Koda Academy Full-Stack Web Development project.

## Technologies

| Technology     | Purpose                               |
| -------------- | ------------------------------------- |
| Go 1.27.1      | Backend programming language          |
| Gin 1.12.0     | HTTP routing and request handling     |
| PostgreSQL 16  | Relational database                   |
| pgx/v5         | PostgreSQL driver and connection pool |
| Redis          | Event data caching                    |
| JWT            | Authentication                        |
| Gin Validator  | Request validation                    |
| Swaggo         | API documentation                     |
| golang-migrate | Database version management           |
| Docker         | Multi-stage container build           |
| Docker Compose | Service orchestration                 |
| Nginx          | Reverse proxy                         |

## Features

### Authentication

- Register and login
- JWT authentication
- Logout and token blacklist
- Forgot and reset password

### Event Management

- Event listing, searching, and pagination
- Event categories
- Event details
- Create, update, and delete events
- Join and leave events
- Saved events
- Event discussions

### Community Management

- Community listing and categories
- Join and leave communities
- Saved communities
- Community member listing
- Community discussions

### User Profile

- Get profile
- Update profile
- Upload avatar

## Project Architecture

The API uses a layered architecture:

```text
cmd/
└── main.go

internal/
├── config/          # PostgreSQL and Redis
├── dto/             # Request and response structures
├── handlers/        # Gin handlers
├── middleware/      # JWT and CORS
├── models/          # Models
├── repositories/    # PostgreSQL queries
├── routes/          # API routing
├── services/        # Business logic
└── utils/           # Shared utilities

db/
├── migrations/      # SQL migrations
└── seeds/           # Initial data

docs/                # Swagger documentation
nginx/               # Reverse proxy configuration
Dockerfile
docker-compose.yml
docker-compose.production.yml
makefile
```

## Getting Started

### Prerequisites

- Git
- Go version compatible with `go.mod`
- Docker and Docker Compose
- Make and golang-migrate CLI for migration commands
- PostgreSQL and Redis, provided by Docker Compose

### Clone Repository

```bash
git clone https://github.com/rlanakomara7/koda-b9-eventhub-backend.git
cd koda-b9-eventhub-backend
```

### Environment Configuration

Create `.env` in the project root.

Example for the Docker Compose environment:

```env
HOST=0.0.0.0
PORT=9000

DBUSER=eventhub
DBPASS=replace_with_database_password
DBHOST=db
DBPORT=5432
DBNAME=koda_eventhub

JWT_SECRET=replace_with_a_secure_random_secret

RDB_HOST=cache
RDB_PORT=6379
RDB_USER=default
RDB_PASS=replace_with_redis_password
```

Replace the sample credentials with your own values.

Never commit real secrets or `.env` files.

### Run with Docker Compose

```bash
docker compose up -d --build
```

Check services:

```bash
docker compose ps
```

View backend logs:

```bash
docker compose logs -f backend
```

Backend API:

http://localhost:8500/api

Swagger UI:

http://localhost:8500/swagger/index.html

Stop without removing database volumes:

```bash
docker compose down
```

## Database Migrations

Migration files are stored in `db/migrations/`.

The Makefile uses `.env.migration` to connect to PostgreSQL.

Available commands:

```bash
make migrate-version
make migrate-up
make migrate-down
make seed
```

**Important:** Verify that `.env.migration` targets the intended database before running these commands.

The Docker database and a local PostgreSQL instance may be different.

Migration files must be tested on a fresh database before production use. Do not force migration versions or reset a database containing important data without verification.

Seeder files are available under `db/seeds/`.

## API Endpoints

Base URL:

http://localhost:8500

### Authentication

| Method | Endpoint                    | Description     |
| ------ | --------------------------- | --------------- |
| POST   | `/api/auth/register`        | Register        |
| POST   | `/api/auth/login`           | Login           |
| POST   | `/api/auth/logout`          | Logout          |
| POST   | `/api/auth/forgot-password` | Forgot password |
| POST   | `/api/auth/reset-password`  | Reset password  |

### Events

| Method | Endpoint                      | Description     |
| ------ | ----------------------------- | --------------- |
| GET    | `/api/events`                 | List events     |
| GET    | `/api/events/:id`             | Event details   |
| GET    | `/api/events/categories`      | Categories      |
| POST   | `/api/events`                 | Create event    |
| PATCH  | `/api/events/:id`             | Update event    |
| DELETE | `/api/events/:id`             | Delete event    |
| POST   | `/api/events/:id/join`        | Join event      |
| DELETE | `/api/events/:id/leave`       | Leave event     |
| GET    | `/api/events/joined`          | Joined events   |
| GET    | `/api/events/saved`           | Saved events    |
| GET    | `/api/events/:id/discussions` | Discussions     |
| POST   | `/api/events/:id/discussions` | Post discussion |

### Communities

| Method | Endpoint                           | Description        |
| ------ | ---------------------------------- | ------------------ |
| GET    | `/api/communities`                 | List communities   |
| GET    | `/api/communities/public`          | Public communities |
| POST   | `/api/communities/:id/join`        | Join community     |
| DELETE | `/api/communities/:id/leave`       | Leave community    |
| GET    | `/api/communities/:id/members`     | Community members  |
| GET    | `/api/communities/saved`           | Saved communities  |
| GET    | `/api/communities/:id/discussions` | Discussions        |
| POST   | `/api/communities/:id/discussions` | Post discussion    |

### User Profile

| Method | Endpoint       | Description    |
| ------ | -------------- | -------------- |
| GET    | `/api/profile` | Get profile    |
| PATCH  | `/api/profile` | Update profile |

Protected endpoints require:

```http
Authorization: Bearer YOUR_TOKEN
```

## Database Transactions

Transactions are implemented in event creation, update, and deletion operations.

Related database changes use `BEGIN`, `COMMIT`, and `ROLLBACK` through pgx transactions.

## Redis Caching

Redis caches event data to reduce repeated database reads.

The event repository uses `GET` and `SET` with a 10-minute TTL, and cache invalidation using `DEL`.

## File Upload

File uploading is supported for event banners and user avatars.

- Accepted extensions: JPG, JPEG, PNG, WEBP
- Maximum file size: 2 MB
- Upload directory: `uploads/`
- Uploading a file is optional where supported

## Gin Validator

Gin request binding validates:

- Required values
- Email formats
- Password length
- Event capacity
- Event formats
- Field length constraints

## CORS

The application uses custom CORS middleware in:

`internal/middleware/cors.middleware.go`

Allowed origins can be configured for the frontend environment.

## Swagger Documentation

Swagger UI:

http://localhost:8500/swagger/index.html

Documentation is generated from Go handler annotations.

After updating annotations, regenerate the Swagger files:

```bash
swag init -g cmd/main.go --parseDependency --parseInternal
```

The committed Swagger specification may need updates for newer endpoints.

## Docker and CI/CD

The backend Dockerfile uses a multi-stage build:

1. Go builder stage compiles the application.
2. Alpine runtime stage executes the compiled binary.

The repository includes Docker Compose and Nginx configuration.

GitHub Actions is used for building and pushing Docker images to GitHub Container Registry (GHCR) after the workflow is configured and verified.

## Testing

```bash
go build ./...
go test ./...
```

## Related Repository

[EventHub Frontend — React](https://github.com/rlanakomara7/EventHub)

## Project

Koda Academy — Full-Stack Web Development.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
