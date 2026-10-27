# EVENTHUB API

[![License](https://img.shields.io/badge/License-MIT-green)](https://opensource.org/license/mit)
[![Go](https://img.shields.io/badge/Go-1.27.1-blue?logo=Go)](https://go.dev/)
![Gin-Gonic](https://img.shields.io/badge/Gin%20Gonic-1.12.0-blue?logo=Gin&logoColor=fff300).
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-16.15-blue?logo=postgresql)](https://www.postgresql.org/)
[![Redis](https://img.shields.io/badge/Redis-9.22.0-blue?logo=Redis)](https://hub.docker.com/_/redishttps://hub.docker.com/_/redis)
[![Swagger](https://img.shields.io/badge/Swagger-1.6.1-blue?logo=swagger)](https://swagger.io)
[![Swagger](https://img.shields.io/badge/Docker-29.8.1-blue?logo=docker)](https://www.docker.com)

<br>

![Logo Koda](https://kodacampmain.github.io/assets-library/img/logo-black.png)

<br>

EventHub API is a REST API backend project developed as part of the Koda Academy Batch 9 backend exercise.

The application provides authentication, event management, community management, user profiles, Redis caching, PostgreSQL persistence, Swagger documentation, and Docker-based deployment.

<br>

## Technologies

### Backend

- Go (Golang)
- Gin Gonic
- REST API
- MVC / Layered Architecture

### Database & Cache

- PostgreSQL
- Redis 7

### Authentication

- JWT (JSON Web Token)
- Token Blacklist

### API Documentation

- Swagger / OpenAPI

### DevOps

- Docker
- Docker Compose
- Nginx Reverse Proxy

### Development Tools

- Git
- GitHub
- Makefile
- Environment Variables (.env)

## Features

- Authentication
  - Register,
  - Login,
  - Logout,
  - Forgot Password,
  - Reset Password.
- User Profile
- Edit User Profile
- Event List
- Event Detail
- Join Event
- Saved Event
- Community List
- Community Detail
- Join Community
- Comment & Testimonial

## Usage Instruction

### DB Setup

1. Create an environment file named `.env` in the root directory.
2. Configure your database credentials:

```env
DB_HOST={YOUR_DB_HOST}
DB_PORT=5432
DB_USER={YOUR_DB_USER}
DB_PASS={YOUR_DB_PASSWORD}
DB_NAME={YOUR_DB_NAME}
```

Add other environment variables required by the application. Do not commit real credentials to GitHub.

### Running the Application

1. Clone this repository:

```bash
git clone https://github.com/rlanakomara7/koda-b9-eventhub-backend.git
cd koda-b9-eventhub-backend
```

2. Configure your `.env` file.

3. Build and run using Docker Compose:

```bash
docker compose up -d --build
```

4. Check running containers:

```bash
docker compose ps
```

5. Access Swagger Documentation:

`http://localhost:8500/swagger/index.html`

To stop the application:

```bash
docker compose down
```

### Running with Frontend

For the integrated production environment, ensure the frontend repository is available at the path configured in Docker Compose.

```bash
docker compose -f docker-compose.yml -f docker-compose.production.yml up -d --build
```

## Routes

### Authentication

| Endpoint             | Method | Description   |
| -------------------- | ------ | ------------- |
| `/api/auth/register` | POST   | Register User |
| `/api/auth/login`    | POST   | Login         |
| `/api/auth/logout`   | POST   | Logout        |

### Events

| Endpoint                | Method | Description  |
| ----------------------- | ------ | ------------ |
| `/api/events`           | GET    | Event List   |
| `/api/events/:id`       | GET    | Event Detail |
| `/api/events`           | POST   | Create Event |
| `/api/events/my`        | GET    | My Events    |
| `/api/events/:id`       | PATCH  | Update Event |
| `/api/events/:id`       | DELETE | Delete Event |
| `/api/events/:id/join`  | POST   | Join Event   |
| `/api/events/:id/leave` | DELETE | Leave Event  |

### Communities

| Endpoint                       | Method | Description        |
| ------------------------------ | ------ | ------------------ |
| `/api/communities`             | GET    | Community List     |
| `/api/communities/public`      | GET    | Public Communities |
| `/api/communities/:id/members` | GET    | Community Members  |
| `/api/communities/:id/join`    | POST   | Join Community     |
| `/api/communities/:id/leave`   | DELETE | Leave Community    |

Some endpoints require JWT authentication.

### Documentation

For complete API documentation, visit `/swagger/index.html`.

## Changelog

| Version    | Description                                                                             |
| ---------- | --------------------------------------------------------------------------------------- |
| Unreleased | Authentication, User Profile, Event Management, Community Management, Community Members |
| Planned    | Event Discussion and Community Discussion                                               |

## How to Contribute

1. Fork this repository.
2. Create a new branch.
3. Commit and push your changes.
4. Open a Pull Request.

## License

This project is licensed under the **MIT License**. Ensure the repository contains a `LICENSE` file.

## Related Project

[EventHub Frontend](https://github.com/rlanakomara7/EventHub)
