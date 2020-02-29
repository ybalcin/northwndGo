package main

import (
	"github.com/gorilla/mux"
	c "github.com/ybalcin/northwndGo/controllers"
	h "github.com/ybalcin/northwndGo/helper"
	"net/http"
	"os"
)

func main() {

	router := mux.NewRouter()
	api := router.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/customers", c.GetAllCustomers).Methods(http.MethodGet)
	api.HandleFunc("/customers/{id}", c.GetCustomer).Methods(http.MethodGet)
	api.HandleFunc("/employees", c.GetAllEmployees).Methods(http.MethodGet)
	api.HandleFunc("/employees/{id}", c.GetEmployee).Methods(http.MethodGet)
	api.HandleFunc("/categories", c.GetAllCategories).Methods(http.MethodGet)
	api.HandleFunc("/categories/{id}", c.GetCategory).Methods(http.MethodGet)
	api.HandleFunc("/employee-territories", c.GetAllEmployeeTerritories).Methods(http.MethodGet)
	api.HandleFunc("/employee-territories/{id}", c.GetEmployeeTerritory).Methods(http.MethodGet)
	api.HandleFunc("/order-details", c.GetAllOrderDetails).Methods(http.MethodGet)
	api.HandleFunc("/order-details/{id}", c.GetOrderDetail).Methods(http.MethodGet)
	api.HandleFunc("/orders", c.GetAllOrders).Methods(http.MethodGet)
	api.HandleFunc("/orders/{id}", c.GetOrder).Methods(http.MethodGet)
	api.HandleFunc("/products", c.GetAllProducts).Methods(http.MethodGet)
	api.HandleFunc("/products/{id}", c.GetProduct).Methods(http.MethodGet)
	api.HandleFunc("/regions", c.GetAllRegion).Methods(http.MethodGet)
	api.HandleFunc("/regions/{id}", c.GetRegion).Methods(http.MethodGet)
	api.HandleFunc("/shippers", c.GetAllShippers).Methods(http.MethodGet)
	api.HandleFunc("/shippers/{id}", c.GetShipper).Methods(http.MethodGet)
	api.HandleFunc("/suppliers", c.GetAllSuppliers).Methods(http.MethodGet)
	api.HandleFunc("/suppliers/{id}", c.GetSupplier).Methods(http.MethodGet)
	api.HandleFunc("/territories", c.GetAllTerritories).Methods(http.MethodGet)
	api.HandleFunc("/territories/{id}", c.GetTerritory).Methods(http.MethodGet)

	port, isExist := os.LookupEnv("PORT")
	if !isExist {
		port = "8000" // localhost
	}

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		h.Log(err)
	}
}
