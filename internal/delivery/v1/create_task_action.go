package v1

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
	"todo/internal/delivery/dto"
	"todo/internal/delivery/response"
	"todo/internal/entity/tasks"
	"todo/internal/usecase"
)

func CreateTaskAction(repo tasks.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var tDto dto.CreateTaskRequest
		err := json.NewDecoder(req.Body).Decode(&tDto)
		if err != nil {
			response.BadRequest(w)
			log.Error(err)
			return
		}
		err = usecase.CreateTask(repo, tDto)
		if err != nil {
			response.InternalServerError(w)
			log.Error(err)
			return
		}
		response.CreateResponse(w, response.Response{Message: "Task Created Successfully."})
	}
}
