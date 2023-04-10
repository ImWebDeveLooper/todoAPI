package repository

import (
	"database/sql"
	"todo/internal/entity/categories"
)

type CategoryRepo struct {
	DB *sql.DB
}
type CategoryTbl struct {
	ID    uint
	Title string
}

func (c CategoryRepo) Marshal(cat categories.Category) (CategoryTbl, error) {
	return CategoryTbl{
		ID:    cat.ID,
		Title: cat.Title,
	}, nil
}

func (c CategoryRepo) UnMarshal(tbl CategoryTbl) (categories.Category, error) {
	return categories.Category{
		ID:    tbl.ID,
		Title: tbl.Title,
	}, nil
}
func (c CategoryRepo) GetCategories() ([]categories.Category, error) {
	query := `SELECT categories.id, categories.title FROM categories`
	selectDB, err := c.DB.Query(query)
	if err != nil {
		return nil, err
	}
	var catsTbl []CategoryTbl
	for selectDB.Next() {
		cat := CategoryTbl{}
		err = selectDB.Scan(&cat.ID, &cat.Title)
		if err != nil {
			return nil, err
		}
		catsTbl = append(catsTbl, cat)
	}
	cats := make([]categories.Category, len(catsTbl))
	for index, category := range catsTbl {
		cats[index], err = c.UnMarshal(category)
		if err != nil {
			return nil, err
		}
	}
	return cats, nil
}

func (c CategoryRepo) GetCategory(cat categories.Category) (categories.Category, error) {
	id := cat.ID
	query := `SELECT * FROM categories WHERE id=?`
	selectDB, err := c.DB.Query(query, id)
	if err != nil {
		return categories.Category{}, err
	}
	sCategory := categories.Category{}
	for selectDB.Next() {
		err = selectDB.Scan(&sCategory.ID, &sCategory.Title)
		if err != nil {
			return categories.Category{}, err
		}
	}
	return sCategory, nil
}

func (c CategoryRepo) CreateCategory(cat categories.Category) error {
	title := cat.Title
	query := `INSERT INTO categories(title) VALUES (?)`
	_, err := c.DB.Exec(query, title)
	if err != nil {
		return err
	}
	return nil
}

func (c CategoryRepo) UpdateCategory(cat categories.Category) error {
	title := cat.Title
	id := cat.ID
	query := `UPDATE categories SET title=? WHERE id=?`
	_, err := c.DB.Exec(query, title, id)
	if err != nil {
		return err
	}
	return nil
}

func (c CategoryRepo) DeleteCategory(cat categories.Category) error {
	query := `DELETE FROM categories WHERE id=?`
	catID := cat.ID
	_, err := c.DB.Exec(query, catID)
	if err != nil {
		return err
	}
	return nil
}
