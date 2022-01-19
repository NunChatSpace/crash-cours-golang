package routes

import (
	"github.com/NunChatSpace/crash-cours-golang/internal/routes/users"
	"github.com/gin-gonic/gin"
)

func BuildRoutes() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "You are reached endpoint.",
		})
	})

	r = users.Build(r)

	return r
}
