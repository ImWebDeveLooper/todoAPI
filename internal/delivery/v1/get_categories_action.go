package v1

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"todo/internal/delivery/response"
	"todo/internal/entity/categories"
	"todo/internal/usecase"
)

// GetCategoriesAction Find all categories
func GetCategoriesAction(repo categories.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		cats, fErr := usecase.GetCategories(repo)
		if fErr != nil {
			log.Error(fErr)
			return
		}
		response.OKResponse(w, response.Response{Data: cats})
	}
}
