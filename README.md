# URL Processor Service

A simple Go HTTP server that processes URLs via an endpoint and provides Swagger API documentation.

---

## Features

- Process URLs through `/process-url` endpoint  
- API documentation available via Swagger UI at `/swagger/index.html`  
- Configurable via `.env` file

---

## Requirements

- Go 1.18+  
- `swaggo/http-swagger` dependency (included via go modules)  

---

## Setup

1. **Clone the repository**

git clone <repository-url>
cd url_processor

2. **Running the Server**
go run .\cmd\api\main.go

3. **Access Swagger UI for API documentation**
http://localhost:8080/swagger/index.html


**Usage Notes**
You can customize the port and default Host value by changing them in .env.