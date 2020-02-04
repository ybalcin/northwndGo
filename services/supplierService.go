package services

import (
	"context"
	"northwndGo/db"
	f "northwndGo/firestore"
	h "northwndGo/helper"
	i "northwndGo/inputs"
	m "northwndGo/models"
)

func GetAllSuppliersFromDb() ([]m.Supplier, error) {
	var e error
	var suppliers []m.Supplier
	db, _ := db.ConnectDb()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM Suppliers")
	if err != nil {
		e = err
	}
	for rows.Next() {
		sp, err := h.RowsToSupplier(rows)
		if err != nil {
			e = err
		}
		suppliers = append(suppliers, sp)
	}
	return suppliers, e
}

func GetAllSuppliers() ([]map[string]interface{}, error) {
	var e error
	suppliers, err := GetAll(i.GetInput{
		Collection: "suppliers",
		Model:      m.Supplier{},
	})
	if err != nil {
		e = err
	}
	return suppliers, e
}

func GetSupplier(id string) (map[string]interface{}, error) {
	var e error
	supplier, err := GetById(i.GetInput{
		Collection: "suppliers",
		Model:      m.Supplier{},
		ID:         id,
	})
	if err != nil {
		e = err
	}
	return supplier, e
}

func WriteToFirestoreSupplier(s m.Supplier) {
	ctx := context.Background()
	client := f.Client()
	defer client.Close()
	_, _, err := client.Collection("suppliers").Add(ctx, m.Supplier{
		SupplierID:   s.SupplierID,
		CompanyName:  s.CompanyName,
		ContactName:  s.ContactName,
		ContactTitle: s.ContactTitle,
		Address:      s.Address,
		City:         s.City,
		Region:       s.Region,
		PostalCode:   s.PostalCode,
		Country:      s.Country,
		Phone:        s.Phone,
		Fax:          s.Fax,
		HomePage:     s.HomePage,
	})
	h.Log(err)
}
