# Go PDF Report Service

A standalone microservice in Go that generates PDF reports for students by consuming the existing Node.js backend API.

## Prerequisites

- Go 1.21 or higher
- Node.js backend running (default: http://localhost:5007)
- PostgreSQL database set up and seeded

## Installation

```bash
cd go-service
go mod tidy
go build -o go-service .
```

## Configuration

Environment variables:

| Variable | Default | Description |
|----------|---------|-------------|
| `GO_SERVICE_PORT` | `8080` | Port for the Go service |
| `NODE_BACKEND_URL` | `http://localhost:5007` | URL of the Node.js backend API |

## Running the Service

```bash
# Using default configuration
./go-service

# With custom configuration
GO_SERVICE_PORT=8081 NODE_BACKEND_URL=http://localhost:5007 ./go-service
```

## API Endpoints

### Health Check
```
GET /health
```
Returns `OK` if the service is running.

### Generate Student Report
```
GET /api/v1/students/:id/report
```

Generates and downloads a PDF report for the specified student.

**Parameters:**
- `id` (path parameter): The student ID

**Response:**
- Content-Type: `application/pdf`
- Content-Disposition: `attachment; filename=student_{id}_report.pdf`

## Example Usage

```bash
# Using curl to download a student report
curl -o student_report.pdf http://localhost:8080/api/v1/students/5/report

# Or open in browser
open http://localhost:8080/api/v1/students/5/report
```

## Architecture

```
go-service/
├── main.go              # Entry point and HTTP server setup
├── handlers/
│   └── student.go       # HTTP request handlers
├── models/
│   └── student.go       # Data models
├── services/
│   ├── api_client.go    # Node.js backend API client
│   └── pdf_generator.go # PDF generation logic
├── go.mod               # Go module definition
└── README.md            # This file
```

## How It Works

1. The service receives a request at `GET /api/v1/students/:id/report`
2. It calls the Node.js backend API at `GET /api/v1/students/:id` to fetch student data
3. The PDF generator creates a formatted PDF report with student information
4. The PDF is returned as a downloadable file

## Testing

1. Ensure the PostgreSQL database is running and seeded
2. Start the Node.js backend (`cd backend && npm start`)
3. Start the Go service (`./go-service`)
4. Make a request to generate a report:
   ```bash
   curl -o report.pdf http://localhost:8080/api/v1/students/5/report
   ```
5. Open `report.pdf` to verify the contents
