package handlers

import (
	"SmartTask/internal/service"

	"go.uber.org/zap"
)

type TaskHandler struct {
	service *service.TaskService
	logger  *zap.Logger
}

func New(service *service.TaskService, logger *zap.Logger) *TaskHandler {
	return &TaskHandler{
		service: service,
		logger:  logger,
	}
}
