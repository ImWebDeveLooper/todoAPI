package usecase

import (
	log "github.com/sirupsen/logrus"
	"todo/internal/delivery/dto"
	"todo/internal/entity/tasks"
)

func DeleteTask(t tasks.Repository, req dto.DeleteTaskRequest) error {
	tsk := tasks.Task{
		ID: req.ID,
	}
	err := t.DeleteTask(tsk)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}
