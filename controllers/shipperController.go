package controllers

import (
	m "github.com/gorilla/mux"
	s "github.com/ybalcin/northwndGo/services"
	u "github.com/ybalcin/northwndGo/utils"
	"net/http"
)

var GetAllShippers = func(w http.ResponseWriter, r *http.Request) {
	shippers, err := s.GetAllShippers()
	if err != nil {
		resp = u.Message(false, err.Error())
	} else {
		resp = u.Message(true, "success")
		resp["data"] = shippers
	}
	u.Respond(w, resp)
}

var GetShipper = func(w http.ResponseWriter, r *http.Request) {
	params := m.Vars(r)
	id := params["id"]
	shipper, err := s.GetShipper(id)
	if err != nil {
		resp = u.Message(false, err.Error())
	} else {
		resp = u.Message(true, "success")
		resp["data"] = shipper
	}
	u.Respond(w, resp)
}
