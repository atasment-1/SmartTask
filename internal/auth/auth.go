package auth

import "go.uber.org/zap"

type AuthService struct {
	logger *zap.Logger
}

func New(logger *zap.Logger) *AuthService {
	return &AuthService{logger: logger}
}
