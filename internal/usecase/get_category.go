package usecase

import (
	"todo/internal/delivery/dto"
	"todo/internal/entity/categories"
)

func GetCategory(c categories.Repository, req dto.GetCategoryRequest) (dto.GetCategoryResponse, error) {
	category := categories.Category{
		ID: req.ID,
	}
	cat, err := c.GetCategory(category)
	if err != nil {
		return dto.GetCategoryResponse{}, err
	}
	return dto.GetCategoryResponse{Category: cat}, nil
}
