package usecase

import (
	log "github.com/sirupsen/logrus"
	"todo/internal/delivery/dto"
	"todo/internal/entity/categories"
)

func DeleteCategory(c categories.Repository, req dto.DeleteCategoryRequest) error {
	cat := categories.Category{
		ID: req.ID,
	}
	err := c.DeleteCategory(cat)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}
