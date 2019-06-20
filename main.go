package main

import (
	"github.com/nodias/go-ApmCommon/middleware"
	"github.com/nodias/go-ApmExam1/router"
	"github.com/urfave/negroni"
)

func main() {
	n := negroni.New(negroni.HandlerFunc(middleware.LoggingMiddleware))
	n.UseHandler(router.NewRouter())
	n.Run(":7001")
}
