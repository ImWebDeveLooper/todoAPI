package usecase

import (
	"todo/internal/delivery/dto"
	"todo/internal/entity/categories"
)

func GetCategories(repo categories.Repository) (dto.GetCategoriesResponse, error) {
	var cats []categories.Category
	cats, err := repo.GetCategories()
	if err != nil {
		return dto.GetCategoriesResponse{}, err
	}
	return dto.GetCategoriesResponse{Categories: cats}, nil
}
