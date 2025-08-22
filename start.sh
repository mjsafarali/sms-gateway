#!/bin/bash
docker-compose up -d mysql-central redis-central nats
docker-compose run --rm api-gateway-migrate
docker-compose up -d