package controllers

import (
	m "github.com/gorilla/mux"
	"net/http"
	s "northwndGo/services"
	u "northwndGo/utils"
)

var GetAllSuppliers = func(w http.ResponseWriter, r *http.Request) {
	suppliers, err := s.GetAllSuppliers()
	if err != nil {
		resp = u.Message(false, err.Error())
	} else {
		resp = u.Message(true, "success")
		resp["data"] = suppliers
	}
	u.Respond(w, resp)
}

var GetSupplier = func(w http.ResponseWriter, r *http.Request) {
	params := m.Vars(r)
	id := params["id"]
	supplier, err := s.GetSupplier(id)
	if err != nil {
		resp = u.Message(false, err.Error())
	} else {
		resp = u.Message(true, "success")
		resp["data"] = supplier
	}
	u.Respond(w, resp)
}
