package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/cr1m5onk1ng/nala_platform_app/api/config"
	"github.com/cr1m5onk1ng/nala_platform_app/api/controllers"
	"github.com/cr1m5onk1ng/nala_platform_app/api/middleware"
	"github.com/cr1m5onk1ng/nala_platform_app/api/routes"
	db "github.com/cr1m5onk1ng/nala_platform_app/db/sqlc"
	"github.com/gofiber/fiber/v2"
)

func StartServerWithGracefulShutdown(a *fiber.App) {
	// Create channel for idle connections.
	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt) // Catch OS signals.
		<-sigint

		// Received an interrupt signal, shutdown.
		if err := a.Shutdown(); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("Oops... Server is not shutting down! Reason: %v", err)
		}

		close(idleConnsClosed)
	}()

	// Run server.
	if err := a.Listen(os.Getenv("SERVER_URL")); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
	}

	<-idleConnsClosed
}

// StartServer func for starting a simple server.
func StartServer(a *fiber.App) {
	// Run server.
	if err := a.Listen(os.Getenv("SERVER_URL")); err != nil {
		log.Printf("Something went wrong while starting the server: %v", err)
	}
}

func main() {
	config := config.FiberConfig()

	app := fiber.New(config)

	// Register basic middleware (log and cors)
	middleware.FiberMiddleware(app)

	database, err := sql.Open("postgres", fmt.Sprintf("dbname=%s password=secret user=root sslmode=disable", "nala"))
	if err != nil {
		panic(err)
	}

	repo := db.NewRepository(database)

	handlers := controllers.NewHandlers(repo)

	// Routes definition.
	routes.SwaggerRoute(app)
	routes.UserRoutes(app, handlers)
	routes.NotFoundRoute(app)

	StartServer(app)

}
