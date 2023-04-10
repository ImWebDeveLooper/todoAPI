package usecase

import (
	"todo/internal/delivery/dto"
	"todo/internal/entity/tasks"
)

func GetTask(t tasks.Repository, req dto.GetTaskRequest) (dto.GetTaskResponse, error) {
	tsk := tasks.Task{
		ID: req.ID,
	}
	task, err := t.GetTask(tsk)
	if err != nil {
		return dto.GetTaskResponse{}, err
	}
	return dto.GetTaskResponse{Task: task}, nil
}
