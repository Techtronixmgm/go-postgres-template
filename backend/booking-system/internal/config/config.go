package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHostName     string
	DBPort         string
	DBUser         string
	DBPassword     string
	DBName         string
	SSLMode        string
	ServerPort     string
	JWTSecret      string
	JWTExpiryHours int
	GinMode        string
}

func Load() (Config, error) {
	paths := []string{
		".env",
		"../.env",
		"../../.env",
	}

	for _, p := range paths {
		if err := godotenv.Load(p); err == nil {
			break
		}
	}
	dbHostName, err := extractEnv("DB_HOST")
	if err != nil {
		return Config{}, err
	}
	dbUser, err := extractEnv("DB_USER")
	if err != nil {
		return Config{}, err
	}

	dbName, err := extractEnv("DB_NAME")
	if err != nil {
		return Config{}, err
	}
	dbPassword, err := extractEnv("DB_PASSWORD")
	if err != nil {
		return Config{}, err
	}
	sslMode, err := extractEnv("DB_SSLMODE")
	if err != nil {
		return Config{}, err
	}

	port, err := extractEnv("PORT")
	if err != nil {
		return Config{}, err
	}

	dbPort, err := extractEnv("DB_PORT")
	if err != nil {
		return Config{}, err
	}

	jwtSecret, err := extractEnv("JWT_SECRET")
	if err != nil {
		return Config{}, err
	}

	jwtExpiryHoursStr, err := extractEnv("JWT_EXPIRY_HOURS")
	if err != nil {
		return Config{}, err
	}

	jwtExpiryHours, err := strconv.Atoi(jwtExpiryHoursStr)
	if err != nil {
		return Config{}, fmt.Errorf("invalid JWT_EXPIRY_HOURS")
	}

	gin_mode, err := extractEnv("GIN_MODE")
	if err != nil {
		return Config{}, err
	}
	return Config{
		DBHostName:     dbHostName,
		DBPort:         dbPort,
		DBUser:         dbUser,
		DBPassword:     dbPassword,
		DBName:         dbName,
		SSLMode:        sslMode,
		ServerPort:     port,
		JWTSecret:      jwtSecret,
		JWTExpiryHours: jwtExpiryHours,
		GinMode:        gin_mode,
	}, nil
}

// Add a Configuration Validation Function
func (c Config) Validate() error {
	if c.DBHostName == "" {
		return errors.New("database host is missing")
	}

	if c.DBPort == "" {
		return errors.New("database port is missing")
	}

	if c.DBUser == "" {
		return errors.New("database user is missing")
	}

	if c.DBPassword == "" {
		return errors.New("database password is missing")
	}

	if c.DBName == "" {
		return errors.New("database name is missing")
	}

	if c.SSLMode == "" {
		return errors.New("database SSL mode is missing")
	}

	if c.ServerPort == "" {
		return errors.New("server port is missing")
	}

	if c.JWTSecret == "" {
		return errors.New("JWT secret is missing")
	}

	if c.JWTExpiryHours <= 0 {
		return errors.New("JWT expiry hours must be greater than 0")
	}

	if c.GinMode == "" {
		return errors.New("Gin mode is missing")
	}

	return nil
}

func extractEnv(key string) (string, error) {
	val := os.Getenv(key)
	if val == "" {
		return "", fmt.Errorf("missing required environment variable: %s", key)
	}
	return val, nil
}
