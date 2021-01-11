package main

import (
	"github.com/gorilla/mux"
	"github.com/pushm0v/gorest/model"
	"github.com/pushm0v/gorest/repository"
	"github.com/pushm0v/gorest/service"
	"log"
	"net/http"
)

func RestRouter() *mux.Router {

	var dbConn, err = NewDBConnection("customer.db")
	if err != nil {
		log.Fatalf("DB Connection error : %v", err)
	}
	dbConn.AutoMigrate(&model.Customer{})
	var custRepository = repository.NewCustomerRepository(dbConn)
	var custService = service.NewCustomerService(custRepository)
	var custHandler = NewCustomerHandler(custService)

	r := mux.NewRouter()
	r.HandleFunc("/api/v1/customers/{id}", custHandler.Get).Methods(http.MethodGet)
	api := r.PathPrefix("/api/v1").Subrouter()

	customerRouter(api, custHandler)
	return r
}

func customerRouter(r *mux.Router, custHandler *CustomerHandler) {
	r.HandleFunc("/customers", custHandler.Post).Methods(http.MethodPost)
	r.HandleFunc("/customers/{id}", custHandler.Put).Methods(http.MethodPut)
	r.HandleFunc("/customers/{id}", custHandler.Delete).Methods(http.MethodDelete)
	r.HandleFunc("/", custHandler.NotFound)
	r.Use(LoggingMiddleware)
}

