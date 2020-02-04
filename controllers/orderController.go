package controllers

import (
	m "github.com/gorilla/mux"
	"net/http"
	s "northwndGo/services"
	u "northwndGo/utils"
)

var GetAllOrders = func(w http.ResponseWriter, r *http.Request) {
	orders, err := s.GetAllOrders()
	if err != nil {
		resp = u.Message(false, err.Error())
	} else {
		resp = u.Message(true, "success")
		resp["data"] = orders
	}
	u.Respond(w, resp)
}

var GetOrder = func(w http.ResponseWriter, r *http.Request) {
	params := m.Vars(r)
	id := params["id"]
	order, err := s.GetOrder(id)
	if err != nil {
		resp = u.Message(false, err.Error())
	} else {
		resp = u.Message(true, "success")
		resp["data"] = order
	}
	u.Respond(w, resp)
}
