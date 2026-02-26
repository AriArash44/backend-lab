package auth

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(h *AuthHandler) *gin.Engine {
	r := gin.Default()
	auth := r.Group("/auth")
	{
		auth.POST("/login", h.login)
		auth.GET("/logout", h.logout)
	}
	return r
}
