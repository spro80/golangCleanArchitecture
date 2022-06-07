package config

import (
	"fmt"
	"os"
)

type ConfigInterface interface {
	Handler() error
}

type ConfigHandler struct {
	port int
	url  string
}

func NewConfig() *ConfigHandler {
	return &ConfigHandler{}
}

func (c ConfigHandler) Handler() (ConfigHandler, error) {

	port := os.Getenv("port")
	url := os.Getenv("url")
	fmt.Printf("Port: %s", port)
	fmt.Printf("Url: %s", url)

	config := ConfigHandler{
		port: 9090,
		url:  "localhost",
	}

	return config, nil
}
