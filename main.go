package main

import (
	"log"
	"net/http"
	"os"

	"github.com/suraj1294/go-web-services-banking/handler"
	"github.com/suraj1294/go-web-services-banking/logger"
)

var defaultPort = "8000"

func main() {

	logger.Info("Starting server...")

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := handler.AppRoutes()

	log.Println("listening on", port)
	log.Fatal(http.ListenAndServe(":"+port, router))

}
