package service_test

import (
	"SmartTask/internal/models"
	"SmartTask/internal/service"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

// Полная mock-реализация интерфейса TaskRepository
type MockTaskRepository struct {
	CreateFunc func(task *models.Task) error
}

func (m *MockTaskRepository) Create(task *models.Task) error {
	return m.CreateFunc(task)
}

func (m *MockTaskRepository) GetByID(id int) (*models.Task, error) {
	return nil, nil
}

func (m *MockTaskRepository) GetByUser(userID int) ([]*models.Task, error) {
	return nil, nil
}

func (m *MockTaskRepository) Update(task *models.Task) error {
	return nil
}

func (m *MockTaskRepository) Delete(id int) error {
	return nil
}

func (m *MockTaskRepository) GetAnalytics(userID int) (*models.TaskAnalytics, error) {
	return nil, nil
}

func TestTaskService(t *testing.T) {
	// Инициализация mock с обработчиком для Create
	mockRepo := &MockTaskRepository{
		CreateFunc: func(task *models.Task) error {
			if task.Title == "" {
				return assert.AnError // Или ваша кастомная ошибка
			}
			return nil
		},
	}

	logger := zap.NewNop()
	serviceInstance := service.NewTaskService(mockRepo, logger)

	t.Run("Create task with empty title", func(t *testing.T) {
		err := serviceInstance.CreateTask(&models.Task{Title: ""}) // Фигурные скобки!
		assert.Error(t, err)
		assert.Equal(t, "title cannot be empty", err.Error())
	})
}
