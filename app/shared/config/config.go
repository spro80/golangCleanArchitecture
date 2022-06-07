package config

import "os"

type ConfigInterface interface {
	Handler() (ConfigHandler, error)
	GetPort() (string, error)
	GetUrl() (string, error)
}

type ConfigHandler struct {
	port string
	url  string
}

func NewConfig() *ConfigHandler {
	return &ConfigHandler{}
}

func (c ConfigHandler) Handler() (ConfigHandler, error) {

	/*
		port := os.Getenv("port")
		url := os.Getenv("url")
		fmt.Printf("Port: %s", port)
		fmt.Printf("Url: %s", url)
	*/
	port, _ := c.GetPort()
	url, _ := c.GetUrl()

	config := ConfigHandler{
		port: port,
		url:  url,
	}

	return config, nil
}

func (c ConfigHandler) GetPort() (string, error) {
	port := os.Getenv("port") //"9091"
	return port, nil
}

func (c ConfigHandler) GetUrl() (string, error) {
	url := os.Getenv("port") //"localhost"
	return url, nil
}
