package repository

import (
	"SmartTask/internal/models"
	"database/sql"
	"time" // Добавьте для работы с time.Duration
)

type PostgresTaskRepository struct {
	db *sql.DB
}

func (r *PostgresTaskRepository) Create(task *models.Task) error {
	query := `INSERT INTO tasks (title, user_id) VALUES ($1, $2) RETURNING id`
	return r.db.QueryRow(query, task.Title, task.UserID).Scan(&task.ID)
}

func (r *PostgresTaskRepository) GetAnalytics(userID int) (*models.TaskAnalytics, error) {
	analytics := &models.TaskAnalytics{
		TasksByCategory: make(map[string]int),
	}

	// Получение общей статистики
	err := r.db.QueryRow(
		`SELECT COUNT(*) FROM tasks WHERE user_id = $1`,
		userID,
	).Scan(&analytics.TotalTasks)
	if err != nil {
		return nil, err
	}

	// Расчет средней продолжительности
	var avgSec float64
	err = r.db.QueryRow(
		`SELECT AVG(EXTRACT(EPOCH FROM (completed_at - created_at)))
         FROM tasks WHERE user_id = $1 AND completed = TRUE`,
		userID,
	).Scan(&avgSec)
	if err != nil {
		return nil, err
	}

	analytics.AvgCompletionTime = time.Duration(avgSec) * time.Second
	return analytics, nil
}
