package application

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"todo/internal/entity/categories"
	"todo/internal/entity/tasks"
	"todo/internal/repository"
)

type App struct {
	DB struct {
		Mysql *sql.DB
	}
	Repository struct {
		TaskRepo     tasks.Repository
		CategoryRepo categories.Repository
	}
	Router *mux.Router
}

// NewApp initialize application
func NewApp() *App {
	app := &App{}
	if err := app.RegisterMysql(); err != nil {
		log.Fatal("Connection Failed!")
	}
	if err := app.RegisterTaskRepo(); err != nil {
		log.Fatal(err)
	}
	if err := app.RegisterCategoryRepo(); err != nil {
		log.Fatal(err)
	}
	if err := app.RegisterRouter(); err != nil {
		log.Fatal(err)
	}
	app.RegisterRoutes()
	return app
}

func (a *App) RegisterMysql() error {
	log.Debug("Database is Connecting...")
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "089223"
	dbName := "crud"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		return err
	}
	a.DB.Mysql = db
	log.Debug("Database Connected Successfully.")
	return nil
}

func (a *App) RegisterTaskRepo() error {
	if a.DB.Mysql == nil {
		return errors.New("cant Connect to Database")
	}
	a.Repository.TaskRepo = repository.TaskRepo{DB: a.DB.Mysql}
	return nil
}

func (a *App) RegisterCategoryRepo() error {
	if a.DB.Mysql == nil {
		return errors.New("cant Connect to Database")
	}
	a.Repository.CategoryRepo = repository.CategoryRepo{DB: a.DB.Mysql}
	return nil
}

func (a *App) RegisterRouter() error {
	a.Router = mux.NewRouter()
	return nil
}

func (a *App) RunRouter() {
	httpListener := &http.Server{
		Addr:    ":3000",
		Handler: a.Router,
	}
	log.Debug("Router is Running...")
	log.Println("Server Started on Port 3000")
	if errL := httpListener.ListenAndServe(); errL != nil {
		log.Fatal(errL)
	}
}
