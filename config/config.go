// Package config reads env if available and initialises config struct
//
// Example:
//  conf := config.New()
package config

import (
	"os"
	"strconv"
	"sync"
)

// GetEnvAsInt looks up for env. If found, returns it else
// returns defaultValue
func GetEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// GetEnvAsInt looks up for env. If found, parses that into int, else
// returns defaultValue
func GetEnvAsInt(key string, defaultValue int) int {
	env := GetEnv(key, strconv.Itoa(defaultValue))
	result, err := strconv.Atoi(env)
	if err != nil {
		panic(err)
	}
	return result
}

// Config struct stores the configuration used througout application
type Config struct {
	ServerPort int    // Port at which the server will listen
	ServerHost string // Host at which the server will listen
}

var conf *Config

// once is a sync.Once used to initialize config once
var once sync.Once

// New initialize config struct to be used throughout application
func New() *Config {
	once.Do(func() {
		conf = &Config{
			ServerPort: GetEnvAsInt("SERVER_PORT", 50051),
			ServerHost: GetEnv("SERVER_HOST", "localhost"),
		}
	})
	return conf
}
