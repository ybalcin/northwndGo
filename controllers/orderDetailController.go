package controllers

import (
	m "github.com/gorilla/mux"
	s "github.com/ybalcin/northwndGo/services"
	u "github.com/ybalcin/northwndGo/utils"
	"net/http"
)

var GetAllOrderDetails = func(w http.ResponseWriter, r *http.Request) {
	o, err := s.GetAllOrderDetails()
	if err != nil {
		resp = u.Message(false, err.Error())
	} else {
		resp = u.Message(true, "success")
		resp["data"] = o
	}
	u.Respond(w, resp)
}

var GetOrderDetail = func(w http.ResponseWriter, r *http.Request) {
	params := m.Vars(r)
	id := params["id"]
	orderDetail, err := s.GetOrderDetail(id)
	if err != nil {
		resp = u.Message(false, err.Error())
	} else {
		resp = u.Message(true, "success")
		resp["data"] = orderDetail
	}
	u.Respond(w, resp)
}
