package main

import (
	"go-ApmCommon/middleware"
	"go-ApmCommon/model"
	"go-ApmExam1/router"

	"github.com/urfave/negroni"
)

var config model.TomlConfig

func init() {
	config.Load("config/%s/config.toml")
}
func main() {
	n := negroni.New(negroni.HandlerFunc(middleware.NewLoggingMiddleware(config.Logpaths.Logpath)))
	n.UseHandler(router.NewRouter())
	n.Run(config.Servers["ApmExam1"].PORT)
}
