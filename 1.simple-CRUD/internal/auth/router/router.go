package router

import (
	"1.simple-CRUD/internal/auth/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter(h *handler.AuthHandler) *gin.Engine {
	r := gin.Default()
	auth := r.Group("/auth")
	{
		auth.POST("/login", h.Login)
		auth.GET("/logout", h.Logout)
	}
	return r
}
