package handler

import (
	"fmt"
	"golb-api/db"
	"log"
	"net"
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

type Handler struct {
	Config    Config `yaml:",inline"`
	DB        db.DB  `yaml:"db" validate:"required"`
	Validator *validator.Validate
}

type Config struct {
	Port string `yaml:"port" validate:"required"`
}

func CreateHandler() (srv Handler) {
	data, err := os.ReadFile("./cfg/cfg.yml")
	if err != nil {
		logrus.Fatalf("Failed to read config file: %v", err)
	}

	err = yaml.Unmarshal(data, &srv)
	if err != nil {
		logrus.Fatalf("Failed to parse config file: %v", err)
	}

	if srv.DB.Config.Password[:6] == "file::" {
		passwordPath := srv.DB.Config.Password[6:]
		correctPath := fmt.Sprintf("./cfg/%s", strings.Replace(passwordPath, "./", "", 1))
		passwordData, err := os.ReadFile(correctPath)
		if err != nil {
			logrus.Fatalf("Failed to read DB password file: %v", err)
		}
		srv.DB.Config.Password = string(passwordData)
	}

	srv.Validator = validator.New()
	err = srv.Validator.Struct(&srv.Config)
	if err != nil {
		logrus.Fatalf("Config validation failed: %v", err)
	}

	srv.DB.Init()

	return
}

func (srv *Handler) CreateGinGroup() {
	router := gin.New()
	router.Use(func(c *gin.Context) {
		start := time.Now()
		c.Next()

		logrus.WithFields(logrus.Fields{
			"status":  c.Writer.Status(),
			"method":  c.Request.Method,
			"path":    c.Request.URL.Path,
			"latency": time.Since(start),
		}).Info("Request processed")
	})

	router.Use(configureCors())

	router.GET("/blogs", srv.HandleGetBlogs)
	router.GET("/blogs/:id", srv.HandleGetBlog)

	router.POST("/blogs", srv.HandleCreateBlog)
	router.POST("/statistics/:blogId", srv.HandleViewBlog)

	runner := fmt.Sprintf("localhost:%s", srv.Config.Port)
	router.Run(runner)
}

func configureCors() gin.HandlerFunc {
	local := fmt.Sprintf("http://%s:3000", getLocalIP())

	return cors.New(cors.Config{
		AllowOrigins:     []string{local, "http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}

func getLocalIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddress := conn.LocalAddr().(*net.UDPAddr)

	return localAddress.IP
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
