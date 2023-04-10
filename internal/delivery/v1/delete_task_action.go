package v1

import (
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"todo/internal/delivery/dto"
	"todo/internal/delivery/response"
	"todo/internal/entity/tasks"
	"todo/internal/usecase"
)

func DeleteTaskAction(repo tasks.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		param := mux.Vars(req)
		tsk := param["taskId"]
		taskID, err := strconv.Atoi(tsk)
		if err != nil {
			response.BadRequest(w)
			log.Error(err)
			return
		}
		var tskDelDto dto.DeleteTaskRequest
		tskDelDto.ID = uint(taskID)
		err = usecase.DeleteTask(repo, tskDelDto)
		if err != nil {
			log.Error(err)
			return
		}
		response.DeleteResponse(w)
	}
}
