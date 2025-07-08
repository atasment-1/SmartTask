package models

import "time"

type TaskAnalytics struct {
	TotalTasks        int            `json:"total_tasks"`
	CompletedTasks    int            `json:"completed_tasks"`
	CompletionRate    float64        `json:"completion_rate"`
	TasksByCategory   map[string]int `json:"tasks_by_category"`
	AvgCompletionTime time.Duration  `json:"avg_completion_time"`
}
