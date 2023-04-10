package application

import (
	"todo/internal/delivery/v1"
)

func (a *App) RegisterRoutes() {
	a.Router.HandleFunc("/tasks", v1.GetTasksAction(a.Repository.TaskRepo)).Methods("GET")
	a.Router.HandleFunc("/tasks/{taskId}", v1.GetTaskAction(a.Repository.TaskRepo)).Methods("GET")
	a.Router.HandleFunc("/tasks", v1.CreateTaskAction(a.Repository.TaskRepo)).Methods("POST")
	a.Router.HandleFunc("/tasks/{taskId}", v1.UpdateTaskAction(a.Repository.TaskRepo)).Methods("PATCH")
	a.Router.HandleFunc("/tasks/{taskId}", v1.DeleteTaskAction(a.Repository.TaskRepo)).Methods("DELETE")
	a.Router.HandleFunc("/categories", v1.GetCategoriesAction(a.Repository.CategoryRepo)).Methods("GET")
	a.Router.HandleFunc("/categories/{categoryId}", v1.GetCategoryAction(a.Repository.CategoryRepo)).Methods("GET")
	a.Router.HandleFunc("/categories", v1.CreateCategoryAction(a.Repository.CategoryRepo)).Methods("POST")
	a.Router.HandleFunc("/categories/{categoryId}", v1.UpdateCategoryAction(a.Repository.CategoryRepo)).Methods("PATCH")
	a.Router.HandleFunc("/categories/{categoryId}", v1.DeleteCategoryAction(a.Repository.CategoryRepo)).Methods("DELETE")
}
