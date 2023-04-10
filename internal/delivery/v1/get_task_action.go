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

func GetTaskAction(repo tasks.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		param := mux.Vars(req)
		tsk := param["taskId"]
		var fTskDto dto.GetTaskRequest
		tskId, err := strconv.Atoi(tsk)
		if err != nil {
			response.NotFound(w)
			log.Error(err)
			return
		}
		fTskDto.ID = uint(tskId)
		gTsk, fErr := usecase.GetTask(repo, fTskDto)
		if err != nil {
			log.Error(fErr)
			return
		}
		response.OKResponse(w, response.Response{Data: gTsk})
	}
}
