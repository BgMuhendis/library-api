# Library-API

## Technologies Used
- [Go](https://go.dev)
- [GORM](https://gorm.io/index.html)
- [Docker](https://www.docker.com)
- [Redis](https://github.com/redis/go-redis)
- [PostgreSQL](https://gorm.io/docs/connecting_to_the_database.html)
- [Swagger](https://github.com/swaggo/swag)

## Using with Docker Compose
```yaml
version: "3.8"

services:
  api:
    container_name: go-server
    build:
      context: .
      dockerfile: build/Dockerfile
    restart: always
    env_file:
      - .env
    environment:
      DATABASE_URL: "host=postgresql user=postgres password=postgres dbname=library sslmode=disable"
    depends_on:
      - postgresql-db
      - redis-cache
    ports:
      - 3000:3000
    networks:
      - library_network
      
  postgresql-db:
    container_name: postgresql
    image: postgres:alpine
    restart: always
    environment: 
      POSTGRES_PASSWORD: postgres
      POSTGRES_USER: postgres
      POSTGRES_DB: library
    ports:
      - 5433:5433
    volumes: 
      - postgres:/var/lib/postgresql/data
    networks:
      - library_network
  
  redis-cache:
    container_name: redis
    image: redis
    restart: always
    ports:
      - 6379:6379
    networks:
      - library_network

 
volumes:
  postgres:

networks:
  library_network:
  
```

## Build and Run

Running the application on docker. To run the application on docker;

```shell
 docker-compose up -d
```
To run the application locally;

First you need to change the redis address to `local`. You need to write `localhost` instead of `redis-cache` in the bottom code snippet written in the book_app go file under the controller folder.

```go
var (
	cacheRedis = cache.NewRedisClient("redis-cache:6379")
)
```

```shell
go run cmd/main.go
```

## API Documentation

<img width="1481" alt="Screenshot 2024-04-27 at 13 42 12" src="https://github.com/BgMuhendis/library-api/assets/34186839/37483601-afee-48e0-9347-af0ad8dbdf41">


