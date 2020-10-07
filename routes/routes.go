package routes

import (
	"Order_Inventory/handler"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Server struct {
	Handler handler.Handler
}

func SetRoutes(r *mux.Router, s *Server) {
	r.HandleFunc("/create/customer", s.Handler.CreateUser).Methods(http.MethodPost)
	r.HandleFunc("/fetch/customer/{id}", s.Handler.FetchUser).Methods(http.MethodGet)
	r.HandleFunc("/fetch/customers", s.Handler.FetchAllUsers).Methods(http.MethodGet)
	r.HandleFunc("/update/customer/{id}", s.Handler.UpdateUser).Methods(http.MethodPost)
	r.HandleFunc("/create/order", s.Handler.CreateOrder).Methods(http.MethodPost)
	r.HandleFunc("/delete/order/{orderId}", s.Handler.CancelOrder).Methods(http.MethodDelete)
	r.HandleFunc("/fetch/customer/{id}/orders", s.Handler.FetchAllOrders).Methods(http.MethodGet)
	err := http.ListenAndServe(":7070", r)
	if err != nil {
		log.Print(err)
	}
}
