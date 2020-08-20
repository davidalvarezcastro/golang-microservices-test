package config

import "os"

const (
	apiGithubAccessToken = "SECRET_GITHUB_ACCESS_TOKEN"
	// LogLevel keeps the level of our logger
	LogLevel      = "info"
	goEnvironment = "GO_ENVIRONMENT"
	prod          = "prod"
)

var (
	githubAccessToken = os.Getenv(apiGithubAccessToken)
)

// GetGithubAccessToken return the api access token
func GetGithubAccessToken() string {
	return githubAccessToken
}

// IsProduction indicates if we are in prod
func IsProduction() bool {
	return os.Getenv(goEnvironment) == prod
}
