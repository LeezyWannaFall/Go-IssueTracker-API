# Go IssueTracker API

REST API for managing issues/tickets.  
Allows creating, retrieving, updating, deleting, and listing issues.  
Mini-product architecture: PostgreSQL, Go, chi router, layered structure.

---

## Technology Stack

- **Language:** Go  
- **Framework/Router:** net/http + chi  
- **Database:** PostgreSQL (Docker container)  
- **Architecture:** Layered (Handler → Service → Repository)  
- **Data Format:** JSON  
- **Additional:** Docker, YAML configuration  

---

## Project Structure
```
Go-IssueTracker-API
├── cmd
│   └── api
│       └── main.go                        # Entry point
├── docker-compose.yml                     # Database container
├── config.yaml
├── go.mod
├── go.sum
├── internal
│   ├── config                             # Configuration functions
│   │   └── config.go
│   ├── handler                            # HTTP handlers
│   │   └── handler.go
│   ├── model                              # Data structures (Issue)
│   │   └── model.go
│   ├── repository                         # Repository interfaces and implementation
│   │   ├── postgres_repo.go
│   │   └── repository.go
│   └── service                            # Business logic
│       └── service.go
├── Makefile
├── migrations                             # SQL migrations
│   ├── 0001_create_issues_table.down.sql
│   └── 0001_create_issues_table.up.sql
└── README.md
```

---

## Setup

### 1. Configure config.yaml

```yaml
server:
  host: "0.0.0.0"
  port: 8080

storage:
  postgres:
    host: "db"
    port: 5432
    database: mydb
    user: task-service
    password: "123456789"
```

- The server will listen at ```http://localhost:server-port-in-config.yaml```

### 2. Start Docker Compose

```bash
docker-compose up --build
```

- Data will be persisted in the volume db_data

### 3. Create the issues table

```bash
docker exec -i issue_tracker_db psql -U task-service -d mydb < migrations/0001_create_issues_table.up.sql
```

Example migration:

```sql
CREATE TABLE IF NOT EXISTS issues (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT,
    status TEXT NOT NULL DEFAULT 'open'
);
```

## API Endpoints

| Method | URL          | Description           |
| ------ | ------------ | --------------------- |
| POST   | /issues      | Create a new issue    |
| GET    | /issues      | List all issues       |
| GET    | /issues/{id} | Get an issue by ID    |
| PUT    | /issues/{id} | Update an issue by ID |
| DELETE | /issues/{id} | Delete an issue by ID |

### Example Requests with curl

- Create an issue:

```bash
curl -X POST http://localhost:8080/issues -H "Content-Type: application/json" -d '{"title": "First issue", "description": "Description"}'
```

- Get an issue by ID

```bash
curl -X GET http://localhost:8080/issues/1
```
or
```bash
curl http://localhost:8080/issues/1
```

- Update an issue:

```bash
curl -X PUT http://localhost:8080/issues/1 -H "Content-Type: application/json" -d '{"title": "Updated issue", "description": "New description", "status": "in_progress"}'
```

- Delete an issue
```bash
curl -X DELETE http://localhost:8080/issues/1
```

- List of issues
```bash
curl http://localhost:8080/issues
```
or
```bash
curl -X GET http://localhost:8080/issues
```
## Notes

- IDs are auto-incremented via PostgreSQL SERIAL. After deleting an issue, new issues will continue incrementing IDs.
