package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/perfectogo/live-chat/api/websocket"
)

func New() *gin.Engine {
	router := gin.New()

	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowCredentials: true,
		AllowHeaders:     []string{"*"},
		AllowMethods:     []string{"*"},
	}))
	router.GET("echo", websocket.WebSocket)

	return router
}
