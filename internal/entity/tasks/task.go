package tasks

type Task struct {
	ID         uint   `json:"taskId"`
	Title      string `json:"taskTitle"`
	IsDone     bool   `json:"status"`
	CategoryID uint   `json:"categoryId"`
}
