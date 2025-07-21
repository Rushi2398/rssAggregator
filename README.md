# RSS Aggregator

A Go-based RSS aggregator with Postgres, Goose migrations, and Docker support.

---

## 🚀 Features

* Fetch RSS feeds concurrently
* Persist feeds & posts to PostgreSQL
* Expose REST API endpoints for users, feeds, and posts
* Database migrations via Goose
* Containerized with Docker & Docker Compose

---

## 📁 Project Structure

```
rssAggregator/
│── main.go
├── internal/
│   ├── auth/
│   └── database/          # DB connection & migration logic
├── sql/                   # Goose migration files
│   └── queries
│   └── schema
├── Dockerfile
├── docker-compose.yml
├── .dockerignore
├── go.mod
├── go.sum
└── README.md
```

---

## 🛠️ Prerequisites

* Go 1.22+
* Docker & Docker Compose
* Goose CLI (`go install github.com/pressly/goose/v3/cmd/goose@latest`)
* PostgreSQL (if running migrations without Docker)

---

## 🧱 Setup & Development

### Clone & Build

```bash
git clone https://github.com/Rushi2398/rssAggregator.git
cd rssAggregator
go mod tidy
```

### Run Migrations

**Locally**:

```bash
goose postgres "postgres://postgres:postgres@localhost:5432/rss_db?sslmode=disable" up
```

---

## 📦 Docker Setup

### docker-compose.yml

```yaml
version: '3.8'
services:
  app:
    build: .
    depends_on: [db]
    ports: ["8080:8080"]
    environment:
      - PORT=8080
      - DB_URL=postgres://postgres:postgres@db:5432/rss_db?sslmode=disable

  db:
    image: postgres:15
    ports: ["5433:5432"]
    volumes: ["postgres-data:/var/lib/postgresql/data"]
    environment:
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD: postgres
      POSTGRES_DB: rss_db

volumes:
  postgres-data:
```

### Build & Run

```bash
docker compose up --build -d
```

---

## ⚙️ Test the App

```bash
curl http://localhost:8080/healthz
```

Check logs:

```bash
docker compose logs -f
```

Check DB:

```bash
psql -h localhost -p 5433 -U postgres -d rss_db
```

---

## 📊 Migrate Local DB Data to Container

```bash
pg_dump -U postgres -d rss_db > dump.sql
docker cp dump.sql <db_container>:/dump.sql
docker exec -it <db_container> psql -U postgres -d rss_db -f /dump.sql
```

---

## ✅ Next Steps

* Add more migration SQLs for new features
* Build out complete HTTP API
* Add Swagger/OpenAPI
* Add CI/CD or Helm charts
