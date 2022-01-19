package users

import (
	"github.com/NunChatSpace/crash-cours-golang/internal/controllers/users"
	"github.com/gin-gonic/gin"
)

func Build(r *gin.Engine) *gin.Engine {
	r.GET("/users", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": users.GetUsers(c.Request.Context()),
		})

	})

	return r
}
