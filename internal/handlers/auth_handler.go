package handlers

import (
	"TideUp/internal/apperror"
	"TideUp/internal/dto"
	"TideUp/internal/services/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *auth.AuthService
}

func NewAuthService(authService *auth.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest
	
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error": "bad request"})
		return
	}

	err := h.authService.Register(req.Email,req.Name,req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"error": "registration failed"})
		return
	}

	c.Status(http.StatusCreated)
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"bad request"})
		return
	}

	token, err := h.authService.Login(req.Email,req.Password)
	if err != nil {
		if err == apperror.ErrBadCredentials {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid password or email"})
		}  else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "login failed"})
		}
		return
	}

	c.JSON(http.StatusOK,gin.H{"token": token})
}

