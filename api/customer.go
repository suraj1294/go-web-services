package api

import (
	"github.com/suraj1294/go-web-services-banking/domain"
	"github.com/suraj1294/go-web-services-banking/errs"
)

var customerService *CustomerService

type CustomerService struct {
	customerRepo *domain.CustomerRepository
}

func (ch CustomerService) GetAllCustomer(status string) ([]domain.Customer, *errs.AppError) {

	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}

	customers, err := ch.customerRepo.FindAll(status)

	return customers, err
}

func (ch CustomerService) GetCustomer(id string) (*domain.Customer, *errs.AppError) {

	customers, err := ch.customerRepo.FindById(id)

	return customers, err
}

func NewCustomerService() *CustomerService {
	customerService = &CustomerService{
		customerRepo: domain.NewCustomerRepository(),
	}

	return customerService
}
