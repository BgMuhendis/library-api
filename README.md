# Library-API

I created the book api with [Fiber](https://docs.gofiber.io/) framework

## Installation

You can visit to [go](https://go.dev/dl/) install Go packages.

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
      - 5432:5432
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

First you need to change the redis address to `local`. You need to write `localhost` instead of `redis-cache` in the bottom code snippet written in the book_app go file under the app folder.

```go
var (
	cacheRedis = cache.NewRedisClient("redis-cache:6379")
)

````
After updating, you can test the application by running the following command

```shell
 go run main.go
```
