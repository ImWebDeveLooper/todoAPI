package usecase

import (
	"todo/internal/delivery/dto"
	"todo/internal/entity/tasks"
)

func CreateTask(t tasks.Repository, req dto.CreateTaskRequest) error {
	task := tasks.Task{
		Title:      req.Title,
		IsDone:     req.IsDone,
		CategoryID: req.CategoryID,
	}
	err := t.CreateTask(task)
	if err != nil {
		return err
	}
	return nil
}
