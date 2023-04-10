package dto

import (
	"todo/internal/entity/tasks"
)

type GetTasksResponse struct {
	Tasks []tasks.Task `json:"tasks"`
}

type CreateTaskRequest struct {
	Title      string `json:"taskTitle"`
	IsDone     bool   `json:"status"`
	CategoryID uint   `json:"categoryId"`
}

type GetTaskRequest struct {
	ID uint `json:"taskId"`
}

type GetTaskResponse struct {
	Task tasks.Task `json:"task"`
}

type UpdateTaskRequest struct {
	ID         uint   `json:"taskId"`
	Title      string `json:"taskTitle"`
	IsDone     bool   `json:"status"`
	CategoryID uint   `json:"categoryId"`
}

type DeleteTaskRequest struct {
	ID uint `json:"taskId"`
}
