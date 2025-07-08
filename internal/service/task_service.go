package service

import (
	"SmartTask/internal/models"
	"SmartTask/internal/repository"

	"go.uber.org/zap"
)

type TaskService struct {
	repo   repository.TaskRepository
	logger *zap.Logger
}

func NewTaskService(repo repository.TaskRepository, logger *zap.Logger) *TaskService {
	return &TaskService{
		repo:   repo,
		logger: logger,
	}
}

func (s *TaskService) CreateTask(task *models.Task) error {
	return s.repo.Create(task)
}
