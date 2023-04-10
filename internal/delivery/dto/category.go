package dto

import (
	"todo/internal/entity/categories"
)

type CreateCategoryRequest struct {
	Title string `json:"categoryTitle"`
}

type UpdateCategoryRequest struct {
	ID    uint   `json:"categoryId"`
	Title string `json:"categoryTitle"`
}

type GetCategoriesResponse struct {
	Categories []categories.Category `json:"categories"`
}

type GetCategoryRequest struct {
	ID uint `json:"categoryId"`
}

type GetCategoryResponse struct {
	Category categories.Category `json:"category"`
}

type DeleteCategoryRequest struct {
	ID uint `json:"categoryId"`
}
