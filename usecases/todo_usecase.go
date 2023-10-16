package usecases

import (
	"boywatz/go-clean/domains"
	"boywatz/go-clean/models"
)

type todoUseCase struct {
	todoRepo domains.TodoRepository
}

// CreateTodo implements domains.TodoUseCase.
func (t *todoUseCase) CreateTodo(todo *models.Todo) (err error) {
	handleError := t.todoRepo.CreateTodo(todo)
	return handleError
}

// GetAllTodo implements domains.TodoUseCase.
func (t *todoUseCase) GetAllTodo() (todos []models.Todo, err error) {
	var todo []models.Todo
	handleError := t.todoRepo.GetAllTodo(&todo)
	return todo, handleError
}

// GetOneTodo implements domains.TodoUseCase.
func (t *todoUseCase) GetOneTodo(todo *models.Todo, id string) (err error) {
	handleError := t.todoRepo.GetOneTodo(todo, id)
	return handleError
}

// UpdateTodo implements domains.TodoUseCase.
func (t *todoUseCase) UpdateTodo(todo *models.Todo, id string) (err error) {
	var todoExist models.Todo
	if errResp := t.todoRepo.GetOneTodo(&todoExist, id); errResp != nil {
		return errResp
	}

	handleError := t.todoRepo.UpdateTodo(&todoExist, id)
	return handleError
}

// DeleteTodo implements domains.TodoUseCase.
func (t *todoUseCase) DeleteTodo(todo *models.Todo, id string) (err error) {
	handleError := t.todoRepo.DeleteTodo(todo, id)
	return handleError
}

func NewTodoUseCase(repo domains.TodoRepository) domains.TodoUseCase {
	return &todoUseCase{todoRepo: repo}
}
