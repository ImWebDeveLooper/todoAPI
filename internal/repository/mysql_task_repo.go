package repository

import (
	"database/sql"
	"todo/internal/entity/tasks"
)

type TaskRepo struct {
	DB *sql.DB
}

func (t TaskRepo) GetTasks() ([]tasks.Task, error) {
	query := `SELECT * FROM tasks`
	selectDB, err := t.DB.Query(query)
	if err != nil {
		return nil, err
	}
	var ts []tasks.Task
	for selectDB.Next() {
		task := tasks.Task{}
		err = selectDB.Scan(&task.ID, &task.Title, &task.IsDone, &task.CategoryID)
		if err != nil {
			return nil, err
		}
		ts = append(ts, task)
	}
	return ts, nil
}

func (t TaskRepo) GetTask(gTask tasks.Task) (tasks.Task, error) {
	id := gTask.ID
	query := `SELECT * FROM tasks WHERE id=?`
	selectDB, err := t.DB.Query(query, id)
	if err != nil {
		return tasks.Task{}, err
	}
	sTask := tasks.Task{}
	for selectDB.Next() {
		err = selectDB.Scan(&sTask.ID, &sTask.Title, &sTask.IsDone, &sTask.CategoryID)
		if err != nil {
			return tasks.Task{}, err
		}
	}
	return sTask, nil
}

func (t TaskRepo) CreateTask(task tasks.Task) error {
	title := task.Title
	status := task.IsDone
	category := task.CategoryID
	query := `INSERT INTO tasks(title,isDone,categoryId) VALUES (?,?,?)`
	_, err := t.DB.Exec(query, title, status, category)
	if err != nil {
		return err
	}
	return nil
}

func (t TaskRepo) UpdateTask(task tasks.Task) error {
	id := task.ID
	title := task.Title
	status := task.IsDone
	catId := task.CategoryID
	query := `UPDATE tasks SET title=?, isdone=? ,categoryId=? WHERE id=?`
	_, err := t.DB.Exec(query, title, status, catId, id)
	if err != nil {
		return err
	}
	return nil
}

func (t TaskRepo) DeleteTask(task tasks.Task) error {
	query := `DELETE FROM tasks WHERE id=?`
	id := task.ID
	_, err := t.DB.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
