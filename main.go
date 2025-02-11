package main

import (
	"auth-sederhana-go-fiber/command"
	"auth-sederhana-go-fiber/config"
	"auth-sederhana-go-fiber/constants"
	"auth-sederhana-go-fiber/routes"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	var env string

	if os.Getenv("GO_ENV") == constants.ENUM_ENV_DEVELOPMENT {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("Error loading env file")
		}
		env = constants.ENUM_ENV_DEVELOPMENT
	} else if os.Getenv("GO_ENV") == constants.ENUM_ENV_PRODUCTION {
		err := godotenv.Load(".env.production")
		if err != nil {
			log.Fatalf("Error loading env file")
		}
		env = constants.ENUM_ENV_PRODUCTION
	} else {
		panic("Invalid GO_ENV")
	}

	fmt.Printf("Environment is %s\n", env)

	db := config.ConnectDB()
	defer config.CloseDBConnection(db)

	if len(os.Args) > 1 {
		flag := command.Commands(db)
		if !flag {
			return
		}
	}

	var serve string

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	switch env {
	case constants.ENUM_ENV_DEVELOPMENT:
		serve = "localhost:" + port

	case constants.ENUM_ENV_PRODUCTION:
		serve = ":" + port
	}

	app := fiber.New()

	routes.SetupRoutes(app, db)

	err := app.Listen(serve)

	if err != nil {
		log.Fatalf("Error when running the server: %v", err)
	}
}
