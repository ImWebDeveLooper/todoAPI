package v1

import (
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"todo/internal/delivery/dto"
	"todo/internal/delivery/response"
	"todo/internal/entity/tasks"
	"todo/internal/usecase"
)

func UpdateTaskAction(repo tasks.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		params := mux.Vars(req)
		id := params["taskId"]
		tskId, err := strconv.Atoi(id)
		if err != nil {
			response.NotFound(w)
			log.Error(err)
			return
		}
		var upTskDto dto.UpdateTaskRequest
		upTskDto.ID = uint(tskId)
		err = json.NewDecoder(req.Body).Decode(&upTskDto)
		if err != nil {
			log.Error(err)
			return
		}
		uErr := usecase.UpdateTask(repo, upTskDto)
		if uErr != nil {
			log.Error(err)
			return
		}
		response.OKResponse(w, response.Response{Message: "Task has been Updated Successfully."})
	}
}
