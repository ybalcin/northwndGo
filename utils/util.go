package utils

import (
	"cloud.google.com/go/firestore"
	"encoding/json"
	h "github.com/ybalcin/northwndGo/helper"
	"net/http"
)

func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func Map(d *firestore.DocumentSnapshot, to interface{}) (interface{}, error) {
	var e error
	bs, err := json.Marshal(d.Data())
	if err != nil {
		e = err
	}
	h.Log(err)
	err = json.Unmarshal(bs, &to)
	if err != nil {
		e = err
	}
	to.(map[string]interface{})["DocID"] = d.Ref.ID
	return to, e
}
