package app

import (
	"SmartTask/internal/config"
	"net/http"

	"go.uber.org/zap"
)

type App struct {
	cfg    *config.Config
	logger *zap.Logger
	server *http.Server
}

func New(cfg *config.Config, logger *zap.Logger) *App {
	return &App{
		cfg:    cfg,
		logger: logger,
	}
}

func (a *App) Run() {
	// Настройка маршрутов
	mux := http.NewServeMux()
	mux.HandleFunc("/", a.handleRoot)
	mux.HandleFunc("/health", a.handleHealthCheck)

	// Создание HTTP сервера
	a.server = &http.Server{
		Addr:    ":" + a.cfg.Server.Port,
		Handler: mux,
	}

	a.logger.Info("Starting server",
		zap.String("port", a.cfg.Server.Port),
	)

	// Запуск сервера
	if err := a.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		a.logger.Fatal("Server failed", zap.Error(err))
	}
}

func (a *App) handleRoot(w http.ResponseWriter, r *http.Request) {
	a.logger.Info("Handling root request")
	w.Write([]byte("Welcome to SmartTask API!"))
}

func (a *App) handleHealthCheck(w http.ResponseWriter, r *http.Request) {
	a.logger.Debug("Health check request")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
