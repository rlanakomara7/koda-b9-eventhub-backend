# EVENTHUB API

[![License](https://img.shields.io/badge/License-MIT-green)](https://opensource.org/license/mit)
[![Go](https://img.shields.io/badge/Go-1.27.1-blue?logo=Go)](https://go.dev/)
![Gin-Gonic](https://img.shields.io/badge/Gin%20Gonic-1.12.0-blue?logo=Gin&logoColor=fff300)
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

<br>

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

### DB setup

Create your environment on the root directory named `.env`

```
DB_HOST={YOUR_DB_HOST}
DB_PORT={YOUR_DB_PORT}
DB_PASS={YOUR_DB_PASS}
DB_NAME={YOUR_DB_NAME}
DB_USER={YOUR_DB_USER}
```

### Running the Application

1. Clone this repostory

```bash
$ git clone url
```

2. Install depedency

```
$ go mod download
```

3. etc.

## Routes

### Features

| Endpoint       | Method | Description    |
| -------------- | ------ | -------------- |
| /auth/register | POST   | Register       |
| /auth/login    | POST   | Login          |
| /auth/logout   | POST   | Logout         |
| /events        | GET    | Event List     |
| /communities   | GET    | Community List |

### Documentation

For complete documentation, visit `/swagger/index.html`

## Changelog

| Version | Desctiption        |
| ------- | ------------------ |
| 1.0.0   | Initial app        |
| 1.1.0   | add authentication |

## How To Contribute

- Fork this repository
- Create your changes
- Pull request

## License

This project is licensed under the MIT License

## Related Project

[Frontend] (https://github.com/rlanakomara7/EventHub.git)
