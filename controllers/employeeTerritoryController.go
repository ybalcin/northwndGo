package controllers

import (
	m "github.com/gorilla/mux"
	"net/http"
	s "northwndGo/services"
	u "northwndGo/utils"
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
