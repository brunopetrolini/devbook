package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// ConnectionString is the connection string to the database
	ConnectionString string = ""
	// Port is the port to run the server on
	Port int = 0
	// SecretKey is the secret key to generate the token
	SecretKey []byte
)

// LoadConfig initializes environment variables
func LoadConfig() {
	var error error

	if error = godotenv.Load(); error != nil {
		log.Fatal(error)
	}

	Port, error = strconv.Atoi(os.Getenv("API_PORT"))
	if error != nil {
		Port = 9000
	}

	ConnectionString = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_DATABASE"),
	)

	SecretKey = []byte(os.Getenv("API_SECRET"))
}
