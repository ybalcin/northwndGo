package services

import (
	"context"
	"github.com/ybalcin/northwndGo/db"
	f "github.com/ybalcin/northwndGo/firestore"
	h "github.com/ybalcin/northwndGo/helper"
	i "github.com/ybalcin/northwndGo/inputs"
	m "github.com/ybalcin/northwndGo/models"
)

func GetAllOrderDetailsFromDb() ([]m.OrderDetail, error) {
	db, _ := db.ConnectDb()
	var e error
	var orderDetails []m.OrderDetail
	defer db.Close()

	rows, err := db.Query("SELECT * FROM [Order Details]")
	if err != nil {
		e = err
	}
	for rows.Next() {
		var o m.OrderDetail
		o, err := h.RowsToOrderDetail(rows)
		if err != nil {
			e = err
		}
		orderDetails = append(orderDetails, o)
	}
	return orderDetails, e
}

func GetAllOrderDetails() ([]map[string]interface{}, error) {
	var e error
	orderDetails, err := GetAll(i.GetInput{
		Collection: "orderDetails",
		Model:      m.OrderDetail{},
	})
	if err != nil {
		e = err
	}
	return orderDetails, e
}

func GetOrderDetail(id string) (map[string]interface{}, error) {
	var e error
	orderDetail, err := GetById(i.GetInput{
		Collection: "orderDetails",
		Model:      m.OrderDetail{},
		ID:         id,
	})
	if err != nil {
		e = err
	}
	return orderDetail, e
}

func WriteToFirestoreOrderDetail(o m.OrderDetail) {
	client := f.Client()
	ctx := context.Background()
	defer client.Close()

	_, _, err = client.Collection("orderDetails").Add(ctx, m.OrderDetail{
		OrderID:   o.OrderID,
		ProductID: o.ProductID,
		UnitPrice: o.UnitPrice,
		Quantity:  o.Quantity,
		Discount:  o.Discount,
	})
	if err != nil {
		h.Log(err)
	}
}
