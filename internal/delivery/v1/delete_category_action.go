package v1

import (
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"todo/internal/delivery/dto"
	"todo/internal/delivery/response"
	"todo/internal/entity/categories"
	"todo/internal/usecase"
)

func DeleteCategoryAction(repo categories.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		param := mux.Vars(req)
		cat := param["categoryId"]
		catID, err := strconv.Atoi(cat)
		if err != nil {
			response.BadRequest(w)
			log.Error(err)
			return
		}
		var catDelDto dto.DeleteCategoryRequest
		catDelDto.ID = uint(catID)
		err = usecase.DeleteCategory(repo, catDelDto)
		if err != nil {
			log.Error(err)
			return
		}
		response.DeleteResponse(w)
	}
}
