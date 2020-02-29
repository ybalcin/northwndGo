package controllers

import (
	m "github.com/gorilla/mux"
	s "github.com/ybalcin/northwndGo/services"
	u "github.com/ybalcin/northwndGo/utils"
	"net/http"
)

var GetAllEmployees = func(w http.ResponseWriter, r *http.Request) {
	employees, err := s.GetAllEmployees()
	if err != nil {
		u.Message(false, err.Error())
	} else {
		resp = u.Message(true, "success")
		resp["data"] = employees
	}
	u.Respond(w, resp)
}

var GetEmployee = func(w http.ResponseWriter, r *http.Request) {
	params := m.Vars(r)
	id := params["id"]
	employee, err := s.GetEmployee(id)
	if err != nil {
		resp = u.Message(false, err.Error())
	} else {
		resp = u.Message(true, "success")
		resp["data"] = employee
	}
	u.Respond(w, resp)
}
