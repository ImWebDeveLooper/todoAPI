package main

import (
	log "github.com/sirupsen/logrus"
	"todo/internal/platform/application"
)

func init() {
	log.SetLevel(log.DebugLevel)
}
func main() {
	app := application.NewApp()
	app.RunRouter()
}
