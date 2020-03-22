package config

import (
	"os"
)

const (
	port = "8989"
)

func GetPort() string {
	p := os.Getenv("HTTP_PORT")
	if p != "" {
		return p
	}
	return port
}
