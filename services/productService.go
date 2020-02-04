package services

import (
	"context"
	"northwndGo/db"
	f "northwndGo/firestore"
	h "northwndGo/helper"
	i "northwndGo/inputs"
	m "northwndGo/models"
)

func GetAllProductsDb() ([]m.Product, error) {
	var e error
	var products []m.Product
	db, _ := db.ConnectDb()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM Products")
	if err != nil {
		e = err
	}
	for rows.Next() {
		p, err := h.RowsToProduct(rows)
		if err != nil {
			e = err
		}
		products = append(products, p)
	}
	return products, e
}

func GetAllProducts() ([]map[string]interface{}, error) {
	var e error
	products, err := GetAll(i.GetInput{
		Collection: "products",
		Model:      m.Product{},
	})
	if err != nil {
		e = err
	}
	return products, e
}

func GetProduct(id string) (map[string]interface{}, error) {
	var e error
	product, err := GetById(i.GetInput{
		Collection: "products",
		Model:      m.Product{},
		ID:         id,
	})
	if err != nil {
		e = err
	}
	return product, e
}

func WriteToFirestoreProduct(p m.Product) {
	client := f.Client()
	defer client.Close()

	ctx := context.Background()
	_, _, err := client.Collection("products").Add(ctx, m.Product{
		ProductID:       p.ProductID,
		ProductName:     p.ProductName,
		SupplierID:      p.SupplierID,
		CategoryID:      p.CategoryID,
		QuantityPerUnit: p.QuantityPerUnit,
		UnitPrice:       p.UnitPrice,
		UnitsInStock:    p.UnitsInStock,
		UnitsOnOrder:    p.UnitsOnOrder,
		ReorderLevel:    p.ReorderLevel,
		Discontinued:    p.Discontinued,
	})
	if err != nil {
		h.Log(err)
	}
}
