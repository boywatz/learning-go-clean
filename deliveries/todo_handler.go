package deliveries

import (
	"boywatz/go-clean/domains"
	"boywatz/go-clean/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type TodoHandler struct {
	todoUseCase domains.TodoUseCase
}

func NewTodoHandler(usecase domains.TodoUseCase) *TodoHandler {
	return &TodoHandler{todoUseCase: usecase}
}

func (t *TodoHandler) GetAllTodo(c *gin.Context) {
	resp, err := t.todoUseCase.GetAllTodo()
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, resp)
	}
}

func (t *TodoHandler) GetOneTodo(c *gin.Context) {
	var todo models.Todo
	id := c.Params.ByName("id")
	err := t.todoUseCase.GetOneTodo(&todo, id)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func (t *TodoHandler) CreateTodo(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBind(&todo); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	err := t.todoUseCase.CreateTodo(&todo)
	if err != nil {
		if ginError := c.AbortWithError(http.StatusInternalServerError, err); ginError != nil {
			logrus.Error(ginError)
		}
	} else {
		c.JSON(http.StatusCreated, todo)
	}
}

func (t *TodoHandler) UpdateTodo(c *gin.Context) {
	id := c.Params.ByName("id")
	var todo models.Todo
	if err := c.ShouldBind(&todo); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	if err := t.todoUseCase.UpdateTodo(&todo, id); err != nil {
		if ginError := c.AbortWithError(http.StatusInternalServerError, err); ginError != nil {
			logrus.Error(ginError)
		}
	} else {
		c.JSON(http.StatusAccepted, todo)
	}
}

func (t *TodoHandler) DeleteTodo(c *gin.Context) {
	id := c.Params.ByName("id")
	var todo models.Todo
	if err := t.todoUseCase.DeleteTodo(&todo, id); err != nil {
		if ginError := c.AbortWithError(http.StatusInternalServerError, err); ginError != nil {
			logrus.Error(ginError)
		}
	} else {
		c.JSON(http.StatusAccepted, gin.H{"id:" + id: "deleted"})
	}
}
