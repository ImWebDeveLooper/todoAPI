package usecase

import (
	"todo/internal/delivery/dto"
	"todo/internal/entity/tasks"
)

func GetTasks(repo tasks.Repository) (dto.GetTasksResponse, error) {
	var tsks []tasks.Task
	tsks, err := repo.GetTasks()
	if err != nil {
		return dto.GetTasksResponse{}, err
	}
	return dto.GetTasksResponse{Tasks: tsks}, nil
}
