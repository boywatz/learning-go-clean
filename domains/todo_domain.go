package domains

import "boywatz/go-clean/models"

type TodoUseCase interface {
	GetAllTodo() (t []models.Todo, err error)
	CreateTodo(t *models.Todo) (err error)
	GetOneTodo(t *models.Todo, id string) (err error)
	UpdateTodo(t *models.Todo, id string) (err error)
	DeleteTodo(t *models.Todo, id string) (err error)
}

type TodoRepository interface {
	GetAllTodo(t *[]models.Todo) (err error)
	CreateTodo(t *models.Todo) (err error)
	GetOneTodo(t *models.Todo, id string) (err error)
	UpdateTodo(t *models.Todo, id string) (err error)
	DeleteTodo(t *models.Todo, id string) (err error)
}
