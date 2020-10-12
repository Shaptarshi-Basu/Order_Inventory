package main

import (
	"Order_Inventory/handler"
	"Order_Inventory/routes"
	"github.com/gorilla/mux"
	"github.com/op/go-logging"
)

func main() {

	r := mux.NewRouter()
	log := logging.MustGetLogger("OrdeInv")
	s := &routes.Server{
		Handler: handler.Handler{*log},
	}
	routes.SetRoutes(r, s)
}
