package main

import (
	"TechNo_Tree_API/handler"
	"TechNo_Tree_API/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	s := &routes.Server{
		Handler: handler.Handler{},
	}
	routes.SetRoutes(r, s)
}

