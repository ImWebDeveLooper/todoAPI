package v1

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"todo/internal/delivery/response"
	"todo/internal/entity/tasks"
	"todo/internal/usecase"
)

func GetTasksAction(repo tasks.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		tsks, err := usecase.GetTasks(repo)
		if err != nil {
			log.Error(err)
			return
		}
		response.OKResponse(w, response.Response{Data: tsks})
	}
}
