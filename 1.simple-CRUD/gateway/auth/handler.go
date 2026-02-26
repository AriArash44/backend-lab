package auth

import (
	"net/http"

	"1.simple-CRUD/common/token"
	"1.simple-CRUD/internal/auth/service"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service *service.AuthService
}

func NewAuthHandler(s *service.AuthService) *AuthHandler {
	return &AuthHandler{service: s}
}

func (h *AuthHandler) login(c *gin.Context) {
	var input loginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	user, err := h.service.Login(input)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	token, _ := token.Generate(user.Username)

	c.SetCookie("session", token, 3600, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "logged in"})
}

func (h *AuthHandler) logout(c *gin.Context) {
	c.SetCookie("session", "", -1, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "logged out"})
}
