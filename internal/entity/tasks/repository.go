package tasks

type Repository interface {
	GetTasks() ([]Task, error)
	GetTask(Task) (Task, error)
	CreateTask(Task) error
	UpdateTask(Task) error
	DeleteTask(Task) error
}
