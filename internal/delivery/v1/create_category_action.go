package v1

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
	"todo/internal/delivery/dto"
	"todo/internal/delivery/response"
	"todo/internal/entity/categories"
	"todo/internal/usecase"
)

func CreateCategoryAction(repo categories.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var dtoCat dto.CreateCategoryRequest
		err := json.NewDecoder(req.Body).Decode(&dtoCat)
		if err != nil {
			response.BadRequest(w)
			log.Error(err)
			return
		}
		err = usecase.CreateCategory(repo, dtoCat)
		if err != nil {
			response.InternalServerError(w)
			log.Error(err)
			return
		}
		response.CreateResponse(w, response.Response{Message: "Category Created Successfully"})
	}
}
