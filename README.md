# Cricket Match Data Microservices

This project contains two microservices written in Go: one for producing cricket match data (Producer) and another for consuming this data (Consumer). The services are containerized using Docker and managed with Docker Compose.

## Project Structure

.
├── docker-compose.yml
├── producer
│ ├── Dockerfile
│ └── main.go
└── consumer
├── Dockerfile
└── main.go


## Prerequisites

- Docker: [Install Docker](https://docs.docker.com/get-docker/)
- Docker Compose: [Install Docker Compose](https://docs.docker.com/compose/install/)
- Postman (optional, for testing): [Download Postman](https://www.postman.com/downloads/)

## Getting Started

### 1. Clone the Repository

```sh
git clone https://github.com/gajare/go-hotstar.git
cd go-hotstar


docker-compose up --build
curl http://localhost:8080/produce
curl http://localhost:8081/consume


output [
  {
    "match_id": "1",
    "score": "TeamA 250/3"
  }
]


docker-compose down
