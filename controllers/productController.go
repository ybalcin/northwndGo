package controllers

import (
	m "github.com/gorilla/mux"
	s "github.com/ybalcin/northwndGo/services"
	u "github.com/ybalcin/northwndGo/utils"
	"net/http"
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
