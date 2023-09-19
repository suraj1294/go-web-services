package handler

import (
	"net/http"

	"github.com/gorilla/mux"
)

func AppRoutes() *mux.Router {

	router := mux.NewRouter()

	router.Headers("content-type", "application/json")
	customerHandler := NewCustomerHandler()
	router.HandleFunc("/customers", customerHandler.GetAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", customerHandler.GetCustomer).Methods(http.MethodGet)

	return router
}
