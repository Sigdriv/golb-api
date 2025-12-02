package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

type Handler struct {
	Config Config `yaml:",inline"`
}

type Config struct {
	Port string `yaml:"port"`
}

func main() {
	handler := &Handler{}

	// Load config from cfg.yml
	data, err := os.ReadFile("./cfg.yml")
	if err != nil {
		logrus.Fatalf("Failed to read config file: %v", err)
	}

	if err := yaml.Unmarshal(data, &handler.Config); err != nil {
		logrus.Fatalf("Failed to parse config file: %v", err)
	}

	log := handler.getLog(nil)
	log.Info("Starting golb-API...")

}

func (*Handler) getLog(c *gin.Context) *logrus.Logger {
	url := c.Request.URL
	logger := logrus.New()

	logger.WithFields(logrus.Fields{
		"method": c.Request.Method,
		"url":    url.String(),
	})

	return logger
}
