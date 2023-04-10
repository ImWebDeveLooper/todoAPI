package categories

type Repository interface {
	GetCategories() ([]Category, error)
	GetCategory(Category) (Category, error)
	CreateCategory(Category) error
	UpdateCategory(Category) error
	DeleteCategory(Category) error
}
