# legalmatch_api

[🇯🇵 日本語](README.md) | 🇺🇸 English

A Go-based REST API for a service that matches users with lawyers for legal consultations. It provides JWT-based authentication, refresh tokens, and management of consultation cases ("agendas").

## Tech Stack

| Category | Technology |
|---|---|
| Language / Framework | Go 1.24, [Gin](https://github.com/gin-gonic/gin) |
| ORM | [GORM](https://gorm.io/) |
| Database | MySQL 8.0 |
| Auth | JWT (access + refresh tokens), bcrypt password hashing |
| Migrations | [goose](https://github.com/pressly/goose) |
| Runtime | Docker Compose |

## Features

- User registration / login / logout
- JWT access token + refresh token authentication
- Role-based access control (user / lawyer, etc.)
- Nickname duplicate check
- Data model and migrations for consultation cases ("agendas")

## Project Structure

```
auth/         Login, token issuance, and refresh logic
config/       DB connection and JWT secret loading
controllers/  User-related request handlers
middlewares/  CORS, logger, and JWT auth middleware
migrations/   Database migrations via goose
models/       GORM models (User, RefreshToken, Agenda)
requests/     Request validation structs
routes/       Route definitions
utils/        JWT parsing and other utilities
```

## Getting Started

### Prerequisites

- Go 1.24+
- Docker / Docker Compose
- [goose](https://github.com/pressly/goose) CLI (for running migrations)

### Steps

1. Clone the repository and set up your environment file

   ```bash
   git clone https://github.com/Twilightcross/legalmatch_api.git
   cd legalmatch_api
   cp .env.example .env
   ```

   Update the values in `.env` (DB credentials, `JWT_SECRET`, etc.) for your environment.

2. Start MySQL via Docker Compose

   ```bash
   docker compose up -d
   ```

3. Run the migrations

   ```bash
   make migrate-up
   ```

4. Start the API server

   ```bash
   go run main.go
   ```

   The server listens on `http://localhost:8080` by default.

## API Endpoints

| Method | Path | Description | Auth |
|---|---|---|---|
| POST | `/api/v1/users/register` | Register a new user | No |
| GET | `/api/v1/users/check-nickname?nickname=xxx` | Check nickname availability | No |
| GET | `/api/v1/users/myinfo` | Get current user's info | Yes |
| POST | `/api/v1/auth/login` | Log in (issues access/refresh tokens) | No |
| POST | `/api/v1/auth/refresh-token` | Reissue an access token | Refresh token |
| POST | `/api/v1/auth/logout` | Revoke a refresh token | Refresh token |
