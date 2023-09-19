package db

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/suraj1294/go-web-services-banking/logger"
)

var localUrl = "postgres://suraj:mysecretpassword@localhost:5433/postgres?sslmode=disable"

type DataBasePg struct {
	url    string
	DbPool *pgxpool.Pool
}

func NewDatabaseConnection() *DataBasePg {

	dbUrl := os.Getenv("DATABASE_URL")

	if dbUrl == "" {
		dbUrl = localUrl
	}

	logger.Info("Initializing DB Connection...")
	databaseInstance := &DataBasePg{
		url: dbUrl,
	}
	DbPool, err := pgxpool.New(context.Background(), dbUrl)

	if err != nil {
		panic(err)
	}

	logger.Info("Connection DB Success!")

	databaseInstance.DbPool = DbPool

	return databaseInstance

}
