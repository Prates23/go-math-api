# Go Math API

This project is a simple RESTful API in Go that supports basic arithmetic operations:
- Sum
- Subtract
- Multiply
- Divide

Each endpoint is versioned (`/api/v1/...`) to support best practices in API design.


## API Usage

### Sum

Request
```bash
GET /api/v1/sum?a=10&b=5
```

Response
```bash
{ "result": 15 }
```

### Subtract

Request
```bash
GET /api/v1/subtract?a=10&b=3
```

Response
```bash
{ "result": 7 }
```

### Multiply

Request
```bash
GET /api/v1/multiply?a=6&b=4
```

Response
```bash
{ "result": 24 }
```

### Divide

Request
```bash
GET /api/v1/divide?a=20&b=4
```

Response
```bash
{ "result": 5 }
```

---

## 🔧 Prerequisites

- Go (version 1.20)
- Docker (optional, for containerized build)

---

## 🚀 Local Development

### 1. Initialize the module

```bash
go mod init go-math-api
go mod tidy
```


### 2. Build Application

```bash
go build -o main .
```

### 3. Run Application

```bash
./main
```

### 4. Run Unit Tests

```bash
go test
```

---

## Containerization

### 1. Build docker image

```bash
docker build -t go-math-api .
```

### 2. Run docker image

```bash
docker run -p 8080:8080 go-math-api
```