# Go Gin API

A RESTful API built with Go and the Gin web framework for managing users. This project includes Swagger documentation and follows REST API best practices.

## Features

- RESTful API endpoints for user management
- Swagger/OpenAPI documentation
- JSON request/response handling
- Input validation
- Error handling
- In-memory data storage (easily replaceable with database)

## Tech Stack

- **Go** 1.24.5
- **Gin** - HTTP web framework
- **Swagger** - API documentation
- **JSON** - Data format

## API Endpoints

### Users

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/v1/users` | Get all users |
| GET | `/api/v1/users/:id` | Get user by ID |
| POST | `/api/v1/users` | Create new user |
| PUT | `/api/v1/users/:id` | Update user |
| DELETE | `/api/v1/users/:id` | Delete user |

### Documentation

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/swagger/*any` | Swagger UI documentation |

## Getting Started

### Prerequisites

- Go 1.24.5 or higher
- Git

### Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd go-gin-api
```

2. Install dependencies:
```bash
go mod tidy
```

3. Run the application:
```bash
go run main.go
```

The server will start on `http://localhost:8080`

### API Documentation

Once the server is running, you can access the Swagger documentation at:
```
http://localhost:8080/swagger/index.html
```

## Usage Examples

### Get all users
```bash
curl -X GET http://localhost:8080/api/v1/users
```

### Get user by ID
```bash
curl -X GET http://localhost:8080/api/v1/users/1
```

### Create a new user
```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com"
  }'
```

### Update a user
```bash
curl -X PUT http://localhost:8080/api/v1/users/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Jane Doe",
    "email": "jane@example.com"
  }'
```

### Delete a user
```bash
curl -X DELETE http://localhost:8080/api/v1/users/1
```

## Project Structure

```
go-gin-api/
├── docs/           # Swagger documentation files
├── handlers/       # HTTP request handlers
├── models/         # Data models and structs
├── routes/         # Route definitions
├── main.go         # Application entry point
├── go.mod          # Go module file
├── go.sum          # Go dependencies checksum
└── README.md       # Project documentation
```

## Data Models

### User
```json
{
  "id": 1,
  "name": "John Doe",
  "email": "john@example.com",
  "created_at": "2025-01-28T10:00:00Z",
  "updated_at": "2025-01-28T10:00:00Z"
}
```

### Create User Request
```json
{
  "name": "John Doe",
  "email": "john@example.com"
}
```

### Update User Request
```json
{
  "name": "Jane Doe",
  "email": "jane@example.com"
}
```

## Development

### Building the application
```bash
go build -o go-gin-api main.go
```

### Running tests
```bash
go test ./...
```

### Generating Swagger docs
If you make changes to the API documentation comments, regenerate the docs:
```bash
swag init
```

## Configuration

The application runs on port 8080 by default. You can modify this in the `main.go` file:

```go
router.Run(":8080") // Change port here
```

## Production Considerations

This is a sample application with in-memory storage. For production use, consider:

- Replace in-memory storage with a proper database (PostgreSQL, MySQL, etc.)
- Add authentication and authorization
- Implement proper logging
- Add rate limiting
- Use environment variables for configuration
- Add comprehensive error handling
- Implement data validation
- Add unit and integration tests
- Set up CI/CD pipeline

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request
