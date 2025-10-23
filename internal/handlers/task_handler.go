package handlers

import (
	"TideUp/internal/dto"
	"TideUp/internal/models"
	"TideUp/internal/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Storage storage.Storage
}

func (h *Handler) AddTask(c *gin.Context) {
	var req dto.CreateTaskRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error": "bad request"})
		return
	}

	task := models.Task{
		Name: req.Name,
		Desc: req.Desc,
		ContextID: req.ContextID,
		Deadline: req.Deadline,
		Completed: false,
	}

	err := h.Storage.AddTask(&task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})
		return
	}
	
	c.JSON(http.StatusOK,gin.H{"response" : "task created"})
}

func (h *Handler) RemoveTask(c* gin.Context) {
	var req dto.DeleteTaskRequest
	
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error": "bad request"})
		return
	}

	err := h.Storage.RemoveTask(req.ID)
	if err !=  nil {
		c.JSON(http.StatusInternalServerError,gin.H{"error": "server error"})
		return
	}

	c.JSON(http.StatusOK,gin.H{"responce":"task successfully removed"})
}

func (h *Handler) ShowAllTasks(c *gin.Context) {
	var req dto.ShowAllTasks

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error": "bad request"})
		return
	}

	resp, err := h.Storage.ShowAllTasks(req.Limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"error": "server error"})
		return
	}

	c.JSON(http.StatusOK,resp)
} 

