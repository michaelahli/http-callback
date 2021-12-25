package main

import (
	"context"
	"fmt"
	"http-callback/helper"
	httpRouter "http-callback/server/http/router"
	"http-callback/server/usecase"
	"http-callback/svcutil/cmd"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi"
	"github.com/rs/cors"
)

func main() {
	config := helper.New()
	config.SetUp()

	var (
		r    = chi.NewRouter()
		host = os.Getenv("APP_HOST")
		bash = cmd.NewTerminal("bash")
	)

	usecase := usecase.UC{
		Helper: config,
		Bash:   bash,
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
