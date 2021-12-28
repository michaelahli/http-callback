package main

import (
	"context"
	"fmt"
	"http-callback/helper"
	httpRouter "http-callback/server/http/router"
	"http-callback/server/usecase"
	"http-callback/svcutil/cmd"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-redis/redis"
	"github.com/rs/cors"
	"github.com/spf13/viper"
)

var config helper.Helper

func init() {
	cfg := viper.NewWithOptions(
		viper.EnvKeyReplacer(
			strings.NewReplacer(".", "_"),
		),
	)

	cfg.SetConfigFile("config")
	cfg.SetConfigType("ini")

	if err := cfg.ReadInConfig(); err != nil {
		log.Fatalf("error loading configuration: %v", err)
	}

	envPath := cfg.GetString("environment.envPath")

	config := helper.New()
	config.SetUp(envPath)
}

func main() {
	var (
		r    = chi.NewRouter()
		host = os.Getenv("APP_HOST")
		bash = cmd.NewTerminal("bash")
	)

	redisClient := &redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	}
	redisPool := redis.NewClient(redisClient)

	pong, err := redisPool.Ping().Result()
	fmt.Println("Redis ping status: "+pong, err)

	usecase := usecase.UC{
		Helper: config,
		Bash:   bash,
		Redis:  redisPool,
	}

	router := httpRouter.New(r, &usecase)
	router.RegisterMiddleware()
	router.RegisterRoutes()

	handler := cors.AllowAll().Handler(r)

	srv := &http.Server{
		Addr:    host,
		Handler: handler,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()
	fmt.Printf("🤖 Server Started. Host: %v \n", host)

	<-done
	fmt.Printf("\tKILL SIGNAL 💀\n🤖 Shutdown signal detected \n")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err := srv.Shutdown(ctx); err != nil {
		panic(err)
	}
	fmt.Println("🤖 Server Stopped Gracefully. \n👋 See you later !")

}
