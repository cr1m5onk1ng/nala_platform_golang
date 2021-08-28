package main

import (
	"database/sql"
	"log"
	"os"
	"os/signal"

	"github.com/cr1m5onk1ng/nala_platform_app/api/config"
	"github.com/cr1m5onk1ng/nala_platform_app/api/controllers"
	"github.com/cr1m5onk1ng/nala_platform_app/api/middleware"
	"github.com/cr1m5onk1ng/nala_platform_app/api/routes"
	repo "github.com/cr1m5onk1ng/nala_platform_app/repository"
	"github.com/cr1m5onk1ng/nala_platform_app/util"
	"github.com/cr1m5onk1ng/nala_platform_app/validation"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func StartServerWithGracefulShutdown(a *fiber.App, serverUrl string) {
	// Create channel for idle connections.
	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt) // Catch OS signals.
		<-sigint

		// Received an interrupt signal, shutdown.
		if err := a.Shutdown(); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("Error while shutting down server: %v", err)
		}
		close(idleConnsClosed)
	}()

	if err := a.Listen(os.Getenv(serverUrl)); err != nil {
		log.Printf("Error while starting server: %v", err)
	}
	<-idleConnsClosed
}

func StartServer(a *fiber.App, serverUrl string) {
	if err := a.Listen(serverUrl); err != nil {
		log.Printf("Something went wrong while starting the server: %v", err)
	}
}

func main() {

	envConfig, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("couldn't load configuration variables")
	}

	fiberConfig := config.FiberConfig()

	app := fiber.New(fiberConfig)

	// Register basic middleware (log and cors)
	middleware.FiberMiddleware(app)

	database, err := sql.Open(envConfig.DB_DRIVER, envConfig.DB_SOURCE)
	if err != nil {
		panic(err)
	}

	repository := repo.NewRepository(database)

	tokenManager, err := validation.NewPasetoTokenManager(envConfig.TOKEN_SYMMETRIC_KEY)
	if err != nil {
		log.Fatal(err.Error())
	}

	handlers, err := controllers.NewHandlers(envConfig, repository, tokenManager)

	if err != nil {
		log.Fatal(err.Error())
	}

	// Routes definition.
	//routes.SwaggerRoute(app)
	routes.ResourceRoutes(app, handlers)
	routes.UserRoutes(app, handlers)
	routes.PostRoutes(app, handlers)
	routes.DiscussionsRoutes(app, handlers)
	routes.NotFoundRoute(app)

	StartServer(app, envConfig.SERVER_URL)

}
