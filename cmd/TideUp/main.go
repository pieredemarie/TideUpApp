package main

import (
	"TideUp/internal/handlers"
	"TideUp/internal/models"
	"TideUp/internal/services/auth"
	"TideUp/internal/services/context"
	"TideUp/internal/services/task"
	"TideUp/internal/storage"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func connectDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("tideup.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = db.AutoMigrate(&models.User{}, &models.Task{}, &models.Context{})
	if err != nil {
		log.Fatal("Failed to migrate database", err)
	}

	return db
}

func registerRoutes(r *gin.Engine, authService *auth.AuthService, taskHandler *handlers.TaskHandler, contextHandler *handlers.ContextHandler) {
	authHandler := handlers.NewAuthService(authService)

	r.POST("/register", authHandler.Register)
	r.POST("/login", authHandler.Login)

	protected := r.Group("/")
	protected.Use(auth.AuthMiddleware())
	{
		protected.POST("/tasks", taskHandler.AddTask)
		protected.GET("/tasks", taskHandler.ShowAllTasks)
		protected.PUT("/tasks/:id", taskHandler.UpdateTask)
		protected.DELETE("/tasks/:id", taskHandler.RemoveTask)

		protected.POST("/contexts", contextHandler.AddContext)
		protected.GET("/contexts", contextHandler.ShowAllContexts)
		protected.PUT("/contexts/:id", contextHandler.EditContext)
		protected.DELETE("/contexts/:id", contextHandler.DeleteContext)
	}
}

func main() {
	godotenv.Load()
	db := connectDB()
	storage := storage.NewStorage(db)

	authService := auth.NewAuthService(storage)
	taskService := task.NewTaskService(storage)
	contextService := context.NewContextService(storage)

	taskHandler := handlers.NewTaskHandler(taskService)
	contextHandler := handlers.NewContextHandler(contextService)
	r := gin.Default()
	registerRoutes(r, authService, taskHandler, contextHandler)
	r.Run(":8080")
}
