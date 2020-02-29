package controllers

import (
	m "github.com/gorilla/mux"
	s "github.com/ybalcin/northwndGo/services"
	u "github.com/ybalcin/northwndGo/utils"
	"net/http"
)

var resp map[string]interface{}

var GetAllCustomers = func(w http.ResponseWriter, r *http.Request) {

	customers, err := s.GetAllCustomers()
	if err != nil {
		u.Message(false, err.Error())
	} else {
		resp = u.Message(true, "success")
		resp["data"] = customers
	}
	u.Respond(w, resp)
}

var GetCustomer = func(w http.ResponseWriter, r *http.Request) {
	params := m.Vars(r)
	id := params["id"]
	customer, err := s.GetCustomer(id)
	if err != nil {
		resp = u.Message(false, err.Error())
	} else {
		resp = u.Message(true, "success")
		resp["data"] = customer
	}
	u.Respond(w, resp)
}
