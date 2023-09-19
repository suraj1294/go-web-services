package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/suraj1294/go-web-services-banking/handler"
	"github.com/suraj1294/go-web-services-banking/logger"
)

func main() {

	logger.Info("Starting server...")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := handler.AppRoutes()

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8000",

		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	//log.Fatal(http.ListenAndServe("localhost:8000", router))

	/** PORT=8000 go run main.go */
	if port != "8080" {
		log.Println("listening on", port)
		log.Fatal(srv.ListenAndServe())
	} else {
		log.Println("listening on", port)
		log.Fatal(http.ListenAndServe(":"+port, router))
	}

}
