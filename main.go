package main

import (
	"context"
	"go-ApmCommon/models"
	"go-ApmCommon/shared/logger"
	"go-ApmExam1/api"
	"go-ApmExam1/router"

	"github.com/urfave/negroni"
)

var config models.TomlConfig

func init() {
	models.Load("config/%s/config.toml")
	config = *models.GetConfig()
	logger.Init()
	api.Init()
}
func main() {
	log := logger.New(context.Background())
	n := negroni.New()
	n.UseHandler(router.NewRouter())
	log.Info("go-ApmExam1 - Server On!")
	n.Run(config.Servers["ApmExam1"].PORT)
}
