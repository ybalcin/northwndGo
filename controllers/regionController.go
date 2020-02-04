package controllers

import (
	m "github.com/gorilla/mux"
	"net/http"
	s "northwndGo/services"
	u "northwndGo/utils"
)

var GetAllRegion = func(w http.ResponseWriter, r *http.Request) {
	regions, err := s.GetAllRegions()
	if err != nil {
		resp = u.Message(false, err.Error())
	} else {
		resp = u.Message(true, "success")
		resp["data"] = regions
	}
	u.Respond(w, resp)
}

var GetRegion = func(w http.ResponseWriter, r *http.Request) {
	params := m.Vars(r)
	id := params["id"]
	region, err := s.GetRegion(id)
	if err != nil {
		resp = u.Message(false, err.Error())
	} else {
		resp = u.Message(true, "success")
		resp["data"] = region
	}
	u.Respond(w, resp)
}
