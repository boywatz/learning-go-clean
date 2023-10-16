package usecases_test

import (
	"boywatz/go-clean/domains/mocks"
	"boywatz/go-clean/models"
	"boywatz/go-clean/usecases"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTodo(t *testing.T) {
	mockRequest := &models.Todo{
		ID:          1,
		Title:       "test",
		Description: "unit test",
	}

	t.Run("success", func(t *testing.T) {
		mockTodoRepo := new(mocks.TodoRepository)
		mockTodoRepo.On("CreateTodo", mockRequest).Return(nil).Once()

		u := usecases.NewTodoUseCase(mockTodoRepo)
		err := u.CreateTodo(mockRequest)

		assert.NoError(t, err)
		assert.Equal(t, mockRequest, mockRequest)
		mockTodoRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockTodoRepo := new(mocks.TodoRepository)
		mockTodoRepo.On("CreateTodo", mockRequest).Return(errors.New("Error")).Once()

		u := usecases.NewTodoUseCase(mockTodoRepo)
		err := u.CreateTodo(mockRequest)

		assert.Error(t, err)
		mockTodoRepo.AssertExpectations((t))
	})
}

func TestGetAllTodo(t *testing.T) {
	var mockResponse []models.Todo

	t.Run("success", func(t *testing.T) {
		mockTodoRepo := new(mocks.TodoRepository)
		mockTodoRepo.On("GetAllTodo", &mockResponse).Return(nil).Once()

		u := usecases.NewTodoUseCase(mockTodoRepo)

		res, err := u.GetAllTodo()
		assert.NoError(t, err)
		assert.Equal(t, res, mockResponse)
		mockTodoRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockTodoRepo := new(mocks.TodoRepository)
		mockTodoRepo.On("GetAllTodo", &mockResponse).Return(errors.New("Error")).Once()

		u := usecases.NewTodoUseCase(mockTodoRepo)

		_, err := u.GetAllTodo()
		assert.Error(t, err)
		mockTodoRepo.AssertExpectations(t)
	})
}
