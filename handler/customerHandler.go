package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/suraj1294/go-web-services-banking/api"
)

type CustomerHandler struct {
	customerService *api.CustomerService
}

func (ch *CustomerHandler) GetAllCustomers(w http.ResponseWriter, r *http.Request) {

	status := r.FormValue("status")

	customers, err := ch.customerService.GetAllCustomer(status)

	if err != nil {
		WriteResponse(w, err.Code, err.AsMessage())
	} else {
		WriteResponse(w, http.StatusOK, customers)
	}

}

func (ch *CustomerHandler) GetCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]

	customer, err := ch.customerService.GetCustomer(id)

	if err != nil {
		WriteResponse(w, err.Code, err.AsMessage())
	} else {
		WriteResponse(w, http.StatusOK, customer)
	}
}

func NewCustomerHandler() *CustomerHandler {
	return &CustomerHandler{api.NewCustomerService()}
}
