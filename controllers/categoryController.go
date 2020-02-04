package controllers

import (
	m "github.com/gorilla/mux"
	"net/http"
	s "northwndGo/services"
	u "northwndGo/utils"
)

var GetAllCategories = func(w http.ResponseWriter, r *http.Request) {
	categories, err := s.GetAllCategories()
	if err != nil {
		resp = u.Message(false, err.Error())
	} else {
		resp = u.Message(true, "success")
		resp["data"] = categories
	}
	u.Respond(w, resp)
}

var GetCategory = func(w http.ResponseWriter, r *http.Request) {
	params := m.Vars(r)
	id := params["id"]
	category, err := s.GetCategory(id)
	if err != nil {
		resp = u.Message(false, err.Error())
	} else {
		resp = u.Message(true, "success")
		resp["data"] = category
	}
	u.Respond(w, resp)
}
