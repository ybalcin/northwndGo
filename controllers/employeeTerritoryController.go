package controllers

import (
	m "github.com/gorilla/mux"
	s "github.com/ybalcin/northwndGo/services"
	u "github.com/ybalcin/northwndGo/utils"
	"net/http"
)

var GetAllEmployeeTerritories = func(w http.ResponseWriter, r *http.Request) {
	et, err := s.GetAllEmployeeTerritories()
	if err != nil {
		resp = u.Message(false, err.Error())
	} else {
		resp = u.Message(true, "success")
		resp["data"] = et
	}
	u.Respond(w, resp)
}

var GetEmployeeTerritory = func(w http.ResponseWriter, r *http.Request) {
	params := m.Vars(r)
	id := params["id"]
	employeeTerritory, err := s.GetEmployeeTerritory(id)
	if err != nil {
		resp = u.Message(false, err.Error())
	} else {
		resp = u.Message(true, "success")
		resp["data"] = employeeTerritory
	}
	u.Respond(w, resp)
}
