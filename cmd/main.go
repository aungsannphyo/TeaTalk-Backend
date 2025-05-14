package main

import (
	"github.com/aungsannphyo/ywartalk/internal/handler"
	"github.com/aungsannphyo/ywartalk/internal/routes"
	"github.com/aungsannphyo/ywartalk/internal/websocket"
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
	hub := websocket.NewHub()
	go hub.Run()
	s.GET("/ws", websocket.WsHandler(hub))
	routes.SetupRoutes(s, handler)
	s.Run(":8080") //localhost:8080
}
