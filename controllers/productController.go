package controllers

import (
	m "github.com/gorilla/mux"
	"net/http"
	s "northwndGo/services"
	u "northwndGo/utils"
)

var GetAllProducts = func(w http.ResponseWriter, r *http.Request) {
	products, err := s.GetAllProducts()
	if err != nil {
		resp = u.Message(false, err.Error())
	} else {
		resp = u.Message(true, "success")
		resp["data"] = products
	}
	u.Respond(w, resp)
}

var GetProduct = func(w http.ResponseWriter, r *http.Request) {
	params := m.Vars(r)
	id := params["id"]
	product, err := s.GetProduct(id)
	if err != nil {
		resp = u.Message(false, err.Error())
	} else {
		resp = u.Message(true, "success")
		resp["data"] = product
	}
	u.Respond(w, resp)
}
