package controllers

import (
	m "github.com/gorilla/mux"
	"net/http"
	s "northwndGo/services"
	u "northwndGo/utils"
)

var GetAllTerritories = func(w http.ResponseWriter, r *http.Request) {
	territories, err := s.GetAllTerritories()
	if err != nil {
		resp = u.Message(false, err.Error())
	} else {
		resp = u.Message(true, "success")
		resp["data"] = territories
	}
	u.Respond(w, resp)
}

var GetTerritory = func(w http.ResponseWriter, r *http.Request) {
	params := m.Vars(r)
	id := params["id"]
	territory, err := s.GetTerritory(id)
	if err != nil {
		resp = u.Message(false, err.Error())
	} else {
		resp = u.Message(true, "success")
		resp["data"] = territory
	}
	u.Respond(w, resp)
}
