# SMS Gateway Monorepo

This is a monorepo with multiple services for SMS management and delivery at scale. It uses Docker Compose with MySQL, Redis, and NATS.

## Quick Start

### Step 1: Run ./start.sh

This will:
- Start MySQL, Redis, and NATS
- Run migrations for api-gateway
- Start all services

### Step 2: Seed database

Run:

```bash
docker exec -i sms_gateway_mysql mysql -u root -ppassword app < ./api-gateway/seeder.sql
```

## Services

- mysql-central → MySQL 8 (database)
- redis-central → Redis 7 (cache, inventory management)
- nats → NATS JetStream (messaging)
- api-gateway → API and input management
- money → Financial service
- money-worker → Background worker
- hermes → SMS delivery service
