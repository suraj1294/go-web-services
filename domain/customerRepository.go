package domain

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/suraj1294/go-web-services-banking/db"
	"github.com/suraj1294/go-web-services-banking/errs"
	"github.com/suraj1294/go-web-services-banking/logger"
)

type Customer struct {
	Id          string    `json:"id" db:"customer_id, pk"`
	Name        string    `json:"name" db:"name"`
	City        string    `json:"city" db:"city"`
	ZipCode     string    `json:"zip_code" db:"zipcode"`
	DateOfBirth time.Time `json:"date_of_birth" db:"date_of_birth"`
	Status      string    `json:"status" db:"status"`
}

type CustomerRepository struct {
	db *db.DataBasePg
}

func (c CustomerRepository) FindAll(status string) ([]Customer, *errs.AppError) {

	ctx := context.Background()

	if err := c.db.DbPool.Ping(ctx); err != nil {
		return nil, errs.NewUnexpectedError("failed to connect DB")
	}

	if status == "" {
		//selectQuery := "SELECT customer_id, name, city, zipCode, date_of_birth, status FROM customers"
		selectQuery := sq.Select("*").From("Customers")
		sql, args, err := selectQuery.ToSql()

		if err != nil {
			logger.Error("error creating query" + err.Error())
			return nil, errs.NewNotFoundError("customer details not found")
		}

		rows, err := c.db.DbPool.Query(ctx, sql, args...)

		if err != nil {
			logger.Error("error fetching details" + err.Error())
			return nil, errs.NewNotFoundError("customer details not found")
		}

		customers, err := pgx.CollectRows(rows, pgx.RowToStructByName[Customer])

		if err != nil {
			logger.Error("error mapping details" + err.Error())
			return nil, errs.NewNotFoundError("customer details not found")
		}
		return customers, nil
	} else {
		selectQuery := "SELECT customer_id, name, city, zipCode, date_of_birth, status FROM customers WHERE status = $1"
		rows, err := c.db.DbPool.Query(ctx, selectQuery, status)
		if err != nil {
			logger.Error("error fetching details" + err.Error())
			return nil, errs.NewNotFoundError("customer details not found")
		}

		customers, err := pgx.CollectRows(rows, pgx.RowToStructByName[Customer])

		if err != nil {
			logger.Error("error mapping details" + err.Error())
			return nil, errs.NewNotFoundError("customer details not found")
		}
		return customers, nil
	}

}

func (c CustomerRepository) FindById(id string) (*Customer, *errs.AppError) {
	ctx := context.Background()

	selectByIdQuery := "SELECT customer_id, name, city, zipCode, date_of_birth, status FROM customers WHERE customer_id = $1"

	var customer Customer
	err := c.db.DbPool.QueryRow(ctx, selectByIdQuery, id).Scan(&customer.Id, &customer.Name, &customer.City, &customer.ZipCode, &customer.DateOfBirth, &customer.Status)
	//err := c.db.QueryRow(ctx, selectByIdQuery, id).Scan(&customer)

	if err != nil {
		logger.Error("error fetching details" + err.Error())
		return nil, errs.NewNotFoundError("customer details not found")
	}

	return &customer, nil

}

func NewCustomerRepository() *CustomerRepository {

	customerRepository := &CustomerRepository{
		db: db.NewDatabaseConnection(),
	}

	return customerRepository
}
