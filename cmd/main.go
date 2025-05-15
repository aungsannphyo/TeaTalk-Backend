package main

import (
	"log"

	"github.com/aungsannphyo/ywartalk/internal/handler"
	"github.com/aungsannphyo/ywartalk/internal/routes"
	"github.com/aungsannphyo/ywartalk/pkg/config"
	"github.com/aungsannphyo/ywartalk/pkg/db"
	database "github.com/aungsannphyo/ywartalk/pkg/db"
	"github.com/gin-gonic/gin"
)

func main() {
	mysqlConfig := config.LoadMySQLConfig()
	database.InitDb(mysqlConfig)
	handler := handler.InitHandler(db.DBInstance)

	s := gin.Default()
	err := s.SetTrustedProxies([]string{
		"127.0.0.1", // localhost
	})
	if err != nil {
		log.Fatal("Failed to set trusted proxies:", err)
	}

	routes.SetupRoutes(s, handler)
	s.Run(":8080") //localhost:8080
}
