package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/ansharw/rest-api/adapters"
	"github.com/ansharw/rest-api/services"
	"github.com/gin-gonic/gin"
)

var env *services.Config

func redisConn() *adapters.RedisDB {
	host := fmt.Sprintf("%s:%s", env.Adapter.Redis.Host, env.Adapter.Redis.Port)
	redisConn := adapters.RedisConnection(host, env.Adapter.Redis.Password, env.Adapter.Redis.DB)
	setRedisConn := &adapters.RedisDB{
		RediConn: redisConn,
	}

	return setRedisConn
}

func main() {
	var err error
	env, err = getConfig()
	if err != nil {
		fmt.Printf("[Main] fail got config [%s]", err)
	}

	services.InitSetEnv(env)

	redis := redisConn()

	svcCockatoo := &services.Cockatoo{}
	svcCockatoo.SetSession(redis)

	bgContext := context.Background()

	// REST
	srvRest := &services.RestHandler{}
	srvRest.RegisterCockatoo(svcCockatoo)

	rest := runHttpServer(srvRest)

	time.Sleep(time.Second * 30)
	ctx, cancel := context.WithTimeout(bgContext, time.Second*time.Duration(env.Service.TimeoutResponseApi))
	if err := rest.Shutdown(ctx); err != nil {
		fmt.Printf("[Main] - failed to gracefully shutdown rest server [%s]\n", err)
	}
	defer cancel()

	if err := redis.CloseRedis(); err != nil {
		fmt.Printf("[Main] - failed to gracefully shutdown redis [%s]\n", err)
	}
}

func runHttpServer(srvRest *services.RestHandler) *http.Server {
	r := gin.New()
	r.Use(gin.Recovery())

	r.POST("/set-session", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(env.Service.TimeoutResponseApi)*time.Millisecond)
		defer cancel()
		srvRest.RestCreateSession(ctx, c.Writer, c.Request)
	})

	r.GET("/version", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "0.0.1",
		})
	})

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})

	srv := &http.Server{
		Addr: ":" + env.Service.Port,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed{
			fmt.Printf("listen: %s\n", err)
		}
	}()

	return srv
}