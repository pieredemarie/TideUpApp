package handlers

import (
	"TideUp/internal/apperror"
	"TideUp/internal/dto"
	"TideUp/internal/models"
	"TideUp/internal/services/context"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ContextHandler struct {
	ContextService context.ContextService
}

func NewContextHandler(ContextService context.ContextService) *ContextHandler {
	return &ContextHandler{
		ContextService: ContextService,
	}
}

func (h *ContextHandler) AddContext(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
		return
	}

	var req dto.ContextRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	context := models.Context{
		UserID:   userID.(int),
		Name:     req.Name,
		Desc:     req.Desc,
		IsHidden: false,
	}

	err := h.ContextService.Create(&context)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		return
	}

	c.JSON(http.StatusCreated, context)
}

func (h *ContextHandler) DeleteContext(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad URL input"})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
		return
	}

	err = h.ContextService.Delete(userID.(int), id)
	if err != nil {
		if errors.Is(apperror.ErrContextNotEmpty, err) {
			c.JSON(http.StatusConflict, gin.H{"error": "context not empty"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		return
	}

	c.Status(http.StatusOK)
}

func (h *ContextHandler) ShowAllContexts(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "10")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 10
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
		return
	}

	contexts, err := h.ContextService.ShowAll(userID.(int), limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		return
	}

	c.JSON(http.StatusOK, contexts)
}

func (h *ContextHandler) EditContext(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad URL input"})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
		return
	}

	var req dto.UpdateContextRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	err = h.ContextService.Edit(userID.(int), id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		return
	}

	c.Status(http.StatusOK)
}
