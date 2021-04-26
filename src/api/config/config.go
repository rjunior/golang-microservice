package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

const (
	apiGithubAccessToken = "SECRET_GITHUB_ACCESS_TOKEN"
	LogLevel             = "info"
	goEnvironment        = "GO_ENVIRONMENT"
	production           = "production"
)

// init is invoked before main()
func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func GetGithubAccessToken() string {
	return os.Getenv(apiGithubAccessToken)
}

func IsProduction() bool {
	return os.Getenv("goEnvironment") == production
}
