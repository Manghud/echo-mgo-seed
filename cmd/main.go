package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/go-playground/validator"
	db "github.com/manghud/echo-mgo-seed/pkg/db"
	v1 "github.com/manghud/echo-mgo-seed/pkg/routes/v1"
	"github.com/manghud/echo-mgo-seed/pkg/utils"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// DefaultPort default running port for server
const DefaultPort = ":7070"

func main() {
	loadEnvironment()
	err := connectToDatabase()
	if err != nil {
		fmt.Println(err)
		return
	}
	port, foundPort := os.LookupEnv("PORT")
	if !foundPort {
		port = DefaultPort
	}
	startServer(port)
}

//DEVELOPMENT contains development environment constant
const DEVELOPMENT string = "development"

func loadEnvironment() {
	environment, found := os.LookupEnv("ENV")
	if !found {
		fmt.Printf("Could not load environment, defaulting to %s\n", DEVELOPMENT)
		environment = DEVELOPMENT
	}
	envFileURI := strings.Join([]string{".env", environment}, ".")
	err := godotenv.Load(envFileURI)
	if err != nil {
		fmt.Printf("Error loading %s file\n", envFileURI)
	}
}

func connectToDatabase() error {
	connectionString, found := os.LookupEnv("MONGODB_URL")
	if !found {
		return errors.New("MONGODB_URL is required")
	}
	dbName, found := os.LookupEnv("DB_NAME")
	if !found {
		return errors.New("DB_NAME is required")
	}
	_, err := db.ConnectToDatabase(connectionString, dbName)
	if err != nil {
		return err
	}
	err = db.CreateDatabaseIndices()
	return err
}

func startServer(port string) {
	server := echo.New()
	server.Pre(middleware.AddTrailingSlash())
	server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAuthorization},
	}))
	server.Use(middleware.Logger())
	server.Use(middleware.Recover())
	server.Validator = &utils.RequestValidator{PayloadValidator: validator.New()}
	v1.RegisterRootRouter(server, "/v1")
	server.Logger.Fatal(server.Start(port))
}
