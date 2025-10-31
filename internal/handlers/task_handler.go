package handlers

import (
	"TideUp/internal/dto"
	"TideUp/internal/models"
	"TideUp/internal/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	Storage storage.Storage
}

func NewTaskHandler(st storage.Storage) *TaskHandler {
	return &TaskHandler{
		Storage: st,
	}
}

func (h *TaskHandler) AddTask(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.AbortWithStatusJSON(http.StatusUnauthorized,gin.H{"error": "user not found"})
		return
	}

	var req dto.TaskRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error": "bad request"})
		return
	}

	task := models.Task{
		UserID: userID.(int),
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
	
	c.JSON(http.StatusOK,task)
}

func (h *TaskHandler) RemoveTask(c* gin.Context) {
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


	err = h.Storage.RemoveTask(userID.(int),id)
	if err !=  nil {
		c.JSON(http.StatusInternalServerError,gin.H{"error": "server error"})
		return
	}

	c.JSON(http.StatusOK,gin.H{"responce":"task successfully removed"})
}

func (h *TaskHandler) ShowAllTasks(c *gin.Context) {
	var req dto.ShowAllTasks

	userID, exists := c.Get("userID")
	if !exists {
		c.AbortWithStatusJSON(http.StatusUnauthorized,gin.H{"error": "user not found"})
		return
	}


	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error": "bad request"})
		return
	}

	resp, err := h.Storage.ShowAllTasks(userID.(int), req.Limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"error": "server error"})
		return
	}

	c.JSON(http.StatusOK,resp)
} 

func (h *TaskHandler) UpdateTask(c *gin.Context) {
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

    var req dto.UpdateTaskRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
        return
    }

    err = h.Storage.UpdateTask(userID.(int),id, req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "db error"})
        return
    }

    c.Status(http.StatusOK)
}
