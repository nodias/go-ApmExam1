package main

import (
	"context"
	"go-ApmCommon/models"
	"go-ApmCommon/shared/logger"
	"go-ApmCommon/shared/middleware"
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
	log := logger.New(context.Background())
	log.Info(config)
}
func main() {
	n := negroni.New(negroni.HandlerFunc(middleware.Logging(config.Logconfig.Logpath)))
	n.UseHandler(router.NewRouter())
	n.Run(config.Servers["ApmExam1"].PORT)
}
