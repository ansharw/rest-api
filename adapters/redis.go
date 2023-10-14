package adapters

import (
	"context"
	"fmt"
	"time"

	"github.com/ansharw/rest-api/api"
	"github.com/go-redis/redis/v8"
)

type RedisDB struct {
	RediConn *redis.Client
}

func RedisConnection(host string, password string, db int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password,
		DB:       db,
	})

	return rdb
}

func (r *RedisDB) CloseRedis() error {
	return r.RediConn.Close()
}

// function for use redis in bussiness logic
func (r *RedisDB) CreateSession(ctx context.Context, table, key string, data *api.Session, ttl time.Duration) error {
	// session:{user_id}
	keys := fmt.Sprintf("%s:%s", table, key)
	redisMap := map[string]interface{}{
		"browser-session": data.BrowserSession,
		"user-agent":      data.UserAgent,
		"user-id":         data.UserId,
		"type": data.Type,
	}

	if errRedis := r.RediConn.HSet(ctx, keys, redisMap).Err(); errRedis != nil {
		return fmt.Errorf("[Redis]CreateSession - failed do HSet [%s]", errRedis.Error())
	}

	if errEx := r.RediConn.Expire(ctx, keys, ttl).Err(); errEx != nil {
		return fmt.Errorf("[Redis]CreateSession - failed set expire [%s]", errEx.Error())
	}

	return nil
}