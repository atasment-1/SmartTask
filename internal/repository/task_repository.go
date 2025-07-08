package repository

import "SmartTask/internal/models" // Добавьте этот импорт

type TaskRepository interface {
	Create(task *models.Task) error
	GetByID(id int) (*models.Task, error)
	GetByUser(userID int) ([]*models.Task, error)
	Update(task *models.Task) error
	Delete(id int) error
	GetAnalytics(userID int) (*models.TaskAnalytics, error)
}
