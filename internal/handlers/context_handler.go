package handlers

import (
	"TideUp/internal/dto"
	"TideUp/internal/models"
	"TideUp/internal/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ContextHandler struct {
	Storage storage.Storage
}

func NewContextHandler(st storage.Storage) *ContextHandler {
	return &ContextHandler{
		Storage: st,
	}
}

func (h *ContextHandler) AddContext(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.AbortWithStatusJSON(http.StatusUnauthorized,gin.H{"error": "user not found"})
		return
	}

	var req dto.ContextRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error": "bad request"})
		return
	}

	context := models.Context{
		UserID: userID.(int),
		Name: req.Name,
		Desc: req.Desc,
		IsHidden: false,
	}

	err := h.Storage.CreateContext(&context)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error" : "server error"})
		return
	}

	c.JSON(http.StatusCreated,context)
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
		c.AbortWithStatusJSON(http.StatusUnauthorized,gin.H{"error": "user not found"})
		return
	}

	err = h.Storage.DeleteContext(userID.(int),id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error" : "server error"})
		return
	}

	c.Status(http.StatusOK)
}

func (h *ContextHandler) ShowAllContexts(c *gin.Context) {
	limitStr := c.DefaultQuery("limit","10")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 10
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.AbortWithStatusJSON(http.StatusUnauthorized,gin.H{"error": "user not found"})
		return
	}
	
	contexts, err := h.Storage.ShowAllContexts(userID.(int),limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error" : "server error"})
		return
	}

	c.JSON(http.StatusOK,contexts)
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
		c.AbortWithStatusJSON(http.StatusUnauthorized,gin.H{"error": "user not found"})
		return
	}

	var req dto.UpdateContextRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error": "bad request"})
		return
	}

	err = h.Storage.EditContext(userID.(int),id,req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error" : "server error"})
		return
	}

	c.Status(http.StatusOK)
}