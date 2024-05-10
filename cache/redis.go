package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
)

type RedisRepository struct {
	client *redis.Client
}

var ctx = context.Background()

func NewRedisClient(address string) *RedisRepository {
	client := redis.NewClient(&redis.Options{
		Addr: address,
		Password: "",
		DB: 0,
 })

 if err := client.Ping(ctx).Err(); err != nil {
		return nil
 }

 return &RedisRepository{
		client: client,
 }

}


func (r *RedisRepository) Get(key string) []byte{
	strCmd  := r.client.Get(ctx,key)

	cacheBytes,err:= strCmd.Bytes()

	if err != nil{
		 return nil
	}

	return cacheBytes
}

func (r *RedisRepository) Set(key string,value []byte)  {

	r.client.Set(ctx,key,value,0)
}

func (r *RedisRepository) Del (key string) {

	r.client.Del(ctx,key)
}