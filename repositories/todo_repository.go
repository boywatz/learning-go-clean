package repositories

import (
	"boywatz/go-clean/domains"
	"boywatz/go-clean/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type todoRepository struct {
	conn *gorm.DB
}

// CreateTodo implements domains.TodoRepository.
func (t *todoRepository) CreateTodo(todo *models.Todo) (err error) {
	if err = t.conn.Create(todo).Error; err != nil {
		return err
	}
	return nil
}

// GetAllTodo implements domains.TodoRepository.
func (t *todoRepository) GetAllTodo(todo *[]models.Todo) (err error) {
	if err = t.conn.Find(todo).Error; err != nil {
		return err
	}
	return nil
}

// GetOneTodo implements domains.TodoRepository.
func (t *todoRepository) GetOneTodo(todo *models.Todo, id string) (err error) {
	if err = t.conn.Where("id = ?", id).First(todo).Error; err != nil {
		return err
	}
	return nil
}

// UpdateTodo implements domains.TodoRepository.
func (t *todoRepository) UpdateTodo(todo *models.Todo, id string) (err error) {
	t.conn.Save(&todo)
	return nil
}

// DeleteTodo implements domains.TodoRepository.
func (t *todoRepository) DeleteTodo(todo *models.Todo, id string) (err error) {
	t.conn.Where("id = ?", id).Delete(todo)
	return nil
}

func NewTodoRepository(conn *gorm.DB) domains.TodoRepository {
	return &todoRepository{conn}
}
