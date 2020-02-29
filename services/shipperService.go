package services

import (
	"context"
	"github.com/ybalcin/northwndGo/db"
	f "github.com/ybalcin/northwndGo/firestore"
	h "github.com/ybalcin/northwndGo/helper"
	i "github.com/ybalcin/northwndGo/inputs"
	m "github.com/ybalcin/northwndGo/models"
)

func GetAllShipperDb() ([]m.Shipper, error) {
	var e error
	var shippers []m.Shipper
	db, _ := db.ConnectDb()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM Shippers")
	if err != nil {
		e = err
	}
	for rows.Next() {
		sh, err := h.RowsToShipper(rows)
		if err != nil {
			e = err
		}
		shippers = append(shippers, sh)
	}
	return shippers, e
}

func GetAllShippers() ([]map[string]interface{}, error) {
	var e error
	shippers, err := GetAll(i.GetInput{
		Collection: "shippers",
		Model:      m.Shipper{},
	})
	if err != nil {
		e = err
	}
	return shippers, e
}

func GetShipper(id string) (map[string]interface{}, error) {
	var e error
	shipper, err := GetById(i.GetInput{
		Collection: "shippers",
		Model:      m.Shipper{},
		ID:         id,
	})
	if err != nil {
		e = err
	}
	return shipper, e
}

func WriteToFirestoreShipper(s m.Shipper) {
	ctx := context.Background()
	client := f.Client()
	defer client.Close()
	_, _, err := client.Collection("shippers").Add(ctx, m.Shipper{
		ShipperID:   s.ShipperID,
		CompanyName: s.CompanyName,
		Phone:       s.Phone,
	})
	h.Log(err)
}
