package main

import (
	"Order_Inventory/handler"
	"Order_Inventory/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	s := &routes.Server{
		Handler: handler.Handler{},
	}
	routes.SetRoutes(r, s)
}

