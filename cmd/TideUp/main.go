package main

import (
	"TideUp/internal/handlers"
	"TideUp/internal/services/auth"
	"TideUp/internal/services/context"
	"TideUp/internal/services/task"
	"TideUp/internal/storage"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func connectDB() *gorm.DB {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")
	dsn := fmt.Sprintf("host=%s port=5432 user=postgres "+
        "password=%s dbname=%s sslmode=disable",host,password,user)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
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
	registerRoutes(r,authService,taskHandler,contextHandler)
	r.Run(":8080")
}