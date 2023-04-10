package usecase

import (
	"todo/internal/delivery/dto"
	"todo/internal/entity/tasks"
)

func UpdateTask(t tasks.Repository, req dto.UpdateTaskRequest) error {
	task := tasks.Task{
		ID:         req.ID,
		Title:      req.Title,
		IsDone:     req.IsDone,
		CategoryID: req.CategoryID,
	}
	err := t.UpdateTask(task)
	if err != nil {
		return err
	}
	return nil
}
