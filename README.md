# Library-API

I created the book api with [Fiber](https://docs.gofiber.io/) framework

## Installation

You can visit to [go](https://go.dev/dl/) install Go packages.

## Build

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
