package router

import (
	"log"
	"net/http"

	"go-ApmExam1/api"

	"github.com/gorilla/mux"
	"go.elastic.co/apm/module/apmgorilla"
)

func NewRouter() *mux.Router {
	return router()
}

func router() (router *mux.Router) {
	router = mux.NewRouter()
	router.HandleFunc("/userInfo/{id}", getUserInfoHandler)
	router.Use(apmgorilla.Middleware())
	return
}

func getUserInfoHandler(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	data, err := api.ApiGetUserInfo(req.Context(), id)
	if err != nil {
		log.Printf("GetUser : %s", err)
		data = []byte(err.Error())
	}
	w.Write(data)
	return
}
