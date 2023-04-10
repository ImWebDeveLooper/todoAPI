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

func GetCategoryAction(repo categories.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		param := mux.Vars(req)
		cat := param["categoryId"]
		var fCatDto dto.GetCategoryRequest
		catId, err := strconv.Atoi(cat)
		if err != nil {
			response.NotFound(w)
			log.Error(err)
			return
		}
		fCatDto.ID = uint(catId)
		gCat, fErr := usecase.GetCategory(repo, fCatDto)
		if fErr != nil {
			log.Error(err)
			return
		}
		response.OKResponse(w, response.Response{Data: gCat})
	}
}
