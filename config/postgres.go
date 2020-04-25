package config

import (
	"os"
)

var (
	DBHost     = os.Getenv("DB_HOST")
	DBName     = os.Getenv("DB_DBNAME")
	DBUser     = os.Getenv("DB_USER")
	DBPassword = os.Getenv("DB_PASSWORD")
	DBSSLMode  = os.Getenv("DB_SSLMODE")
)
