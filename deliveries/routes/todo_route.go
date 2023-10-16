package routes

import (
	"boywatz/go-clean/databases"
	"boywatz/go-clean/deliveries"
	"boywatz/go-clean/repositories"
	"boywatz/go-clean/usecases"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	todoRepo := repositories.NewTodoRepository(databases.DB)
	todoUseCase := usecases.NewTodoUseCase(todoRepo)
	todoHandler := deliveries.NewTodoHandler(todoUseCase)

	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("todo", todoHandler.GetAllTodo)
		v1.POST("todo", todoHandler.CreateTodo)
		v1.GET("todo/:id", todoHandler.GetOneTodo)
		v1.PUT("todo/:id", todoHandler.UpdateTodo)
		v1.DELETE("todo/:id", todoHandler.DeleteTodo)
	}
	return r
}
