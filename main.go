package main

import (
	"barebone-go-crud/src/configs"
	"barebone-go-crud/src/handler"
	"barebone-go-crud/src/repositories"
	"barebone-go-crud/src/router"
	"barebone-go-crud/src/services"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {

	cfg, err := configs.LoadConfig()
	if err != nil {
		log.Fatal("Error loading config : %v", err)
	}

	defer cfg.DB.Close()

	// dependency injection
	userRepo := repositories.NewUserRepository(cfg.DB)
	userService := services.NewUserService(userRepo)
	userHandler := handler.NewHandleUser(userService)

	// setup routing
	mux := router.NewRouter(userHandler)

	srv := &http.Server{
		Addr:    ":" + cfg.PORT,
		Handler: mux,
	}

	// Jalankan server secara goroutine
	go func() {
		log.Printf("Server berjalan di port %s", cfg.PORT)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error menjalankan server: %v", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Server shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown error: %v", err)
	}
	log.Println("Server keluar")
}
