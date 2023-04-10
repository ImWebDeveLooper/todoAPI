package v1

import (
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"todo/internal/delivery/dto"
	"todo/internal/delivery/response"
	"todo/internal/entity/categories"
	"todo/internal/usecase"
)

func UpdateCategoryAction(repo categories.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		param := mux.Vars(req)
		cat := param["categoryId"]
		catId, err := strconv.Atoi(cat)
		if err != nil {
			response.NotFound(w)
			log.Error(err)
			return
		}
		var upCatDto dto.UpdateCategoryRequest
		upCatDto.ID = uint(catId)
		err = json.NewDecoder(req.Body).Decode(&upCatDto)
		if err != nil {
			log.Error(err)
			return
		}
		uErr := usecase.UpdateCategory(repo, upCatDto)
		if uErr != nil {
			log.Error(err)
			return
		}
		response.OKResponse(w, response.Response{Message: "Category has been Updated Successfully."})
	}
}
