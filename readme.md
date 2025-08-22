# SMS Gateway Monorepo

This is a monorepo with multiple services for SMS management and delivery at scale. The project is built using Go microservices architecture with Docker Compose for orchestration, utilizing MySQL for data persistence, Redis for caching and inventory management, and NATS JetStream for reliable messaging between services.

## Quick Start

### Step 1: Start the Services

Run the startup script to initialize the entire system:

```bash
./start.sh
```

This command will:
- Start MySQL, Redis, and NATS containers
- Run database migrations for api-gateway
- Start all microservices

### Step 2: Seed Database

After the services are running, seed the database with initial data:

```bash
docker exec -i sms_gateway_mysql mysql -u root -ppassword app < ./api-gateway/seeders.sql
```

## Services

| Service | Description |
|---------|-------------|
| **mysql-central** | MySQL 8 database for persistent data storage |
| **redis-central** | Redis 7 for caching and inventory management |
| **nats** | NATS JetStream for reliable messaging between services |
| **api-gateway** | Main API service for input management and external communication |
| **money** | Financial service handling transactions and wallet operations |
| **money-worker** | Background worker for processing financial tasks |
| **hermes** | SMS delivery service responsible for message dispatch |

## Architecture

The system follows a microservices architecture pattern where each service has its own responsibility:

- **API Gateway** serves as the entry point for all external requests
- **Money Service** handles all financial operations and wallet management
- **Hermes** manages SMS delivery and provider integrations
- **NATS** ensures reliable communication between services
- **Redis** provides fast caching and real-time inventory tracking
- **MySQL** stores all persistent data with proper migrations
