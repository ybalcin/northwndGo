package services

import (
	"context"
	"northwndGo/db"
	f "northwndGo/firestore"
	h "northwndGo/helper"
	i "northwndGo/inputs"
	m "northwndGo/models"
)

func GetAllOrdersDb() ([]m.Order, error) {
	db, _ := db.ConnectDb()
	defer db.Close()

	var e error
	var orders []m.Order
	rows, err := db.Query("SELECT * FROM Orders")
	if err != nil {
		e = err
	}
	for rows.Next() {
		o, err := h.RowsToOrder(rows)
		if err != nil {
			e = err
		}
		orders = append(orders, o)
	}
	return orders, e
}

func GetAllOrders() ([]map[string]interface{}, error) {
	var e error
	orders, err := GetAll(i.GetInput{
		Collection: "orders",
		Model:      m.Order{},
	})
	if err != nil {
		e = err
	}
	return orders, e
}

func GetOrder(id string) (map[string]interface{}, error) {
	var e error
	order, err := GetById(i.GetInput{
		Collection: "orders",
		Model:      m.Order{},
		ID:         id,
	})
	if err != nil {
		e = err
	}
	return order, e
}

func WriteToFirestoreOrder(o m.Order) {
	client := f.Client()
	defer client.Close()

	ctx := context.Background()
	_, _, err := client.Collection("orders").Add(ctx, m.Order{
		OrderID:        o.OrderID,
		CustomerID:     o.CustomerID,
		EmployeeID:     o.EmployeeID,
		OrderDate:      o.OrderDate,
		RequiredDate:   o.RequiredDate,
		ShippedDate:    o.ShippedDate,
		ShipVia:        o.ShipVia,
		Freight:        o.Freight,
		ShipName:       o.ShipName,
		ShipAddress:    o.ShipAddress,
		ShipCity:       o.ShipCity,
		ShipRegion:     o.ShipRegion,
		ShipPostalCode: o.ShipPostalCode,
		ShipCountry:    o.ShipCountry,
	})
	if err != nil {
		h.Log(err)
	}
}
