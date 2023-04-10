package usecase

import (
	"todo/internal/delivery/dto"
	"todo/internal/entity/categories"
)

func UpdateCategory(c categories.Repository, req dto.UpdateCategoryRequest) error {
	category := categories.Category{
		ID:    req.ID,
		Title: req.Title,
	}
	err := c.UpdateCategory(category)
	if err != nil {
		return err
	}
	return nil
}
