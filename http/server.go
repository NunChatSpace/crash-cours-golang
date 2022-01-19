package http

import (
	"github.com/NunChatSpace/crash-cours-golang/internal/routes"
	"github.com/gin-gonic/gin"
)

func Build() *gin.Engine {
	r := routes.BuildRoutes()

	return r
}
