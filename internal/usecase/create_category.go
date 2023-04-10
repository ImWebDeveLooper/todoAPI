package usecase

import (
	"todo/internal/delivery/dto"
	"todo/internal/entity/categories"
)

func CreateCategory(c categories.Repository, req dto.CreateCategoryRequest) error {
	category := categories.Category{
		Title: req.Title,
	}
	err := c.CreateCategory(category)
	if err != nil {
		return err
	}
	return nil
}
