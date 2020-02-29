package services

import (
	"context"
	"database/sql"
	"github.com/ybalcin/northwndGo/db"
	f "github.com/ybalcin/northwndGo/firestore"
	h "github.com/ybalcin/northwndGo/helper"
	i "github.com/ybalcin/northwndGo/inputs"
	m "github.com/ybalcin/northwndGo/models"
	"os"
)

var err error
var rows *sql.Rows
var Db *sql.DB
var ctx context.Context
var customers []m.Customer

func GetAllCustomersFromDb() ([]m.Customer, error) {
	Db, ctx = db.ConnectDb()

	rows, err = Db.Query("SELECT * FROM Customers")
	defer Db.Close()

	for rows.Next() {
		var customer m.Customer
		customer, err = h.RowsToCustomer(rows)
		customers = append(customers, customer)
	}
	return customers, err
}

func GetByIdFromDb(id string) (m.Customer, error) {
	Db, ctx = db.ConnectDb()
	var customer m.Customer
	param := "'" + id + "'"
	row := Db.QueryRow("SELECT * FROM Customers WHERE CustomerID = " + param)
	customer, err = h.RowToCustomer(row)
	return customer, err
}

func GetAllCustomers() ([]map[string]interface{}, error) {
	var e error
	customers, err := GetAll(i.GetInput{
		Collection: "customers",
		Model:      m.Customer{},
	})
	if err != nil {
		e = err
	}
	return customers, e
}

func GetCustomer(id string) (map[string]interface{}, error) {
	var e error
	customer, err := GetById(i.GetInput{
		Collection: "customers",
		Model:      m.Customer{},
		ID:         id,
	})
	if err != nil {
		e = err
	}
	return customer, e
}

func WriteToFirebaseCustomer(c m.Customer) {
	client := f.Client()
	defer client.Close()

	_, _, err := client.Collection("customers").Add(ctx, m.Customer{
		CustomerID:   c.CustomerID,
		CompanyName:  c.CompanyName,
		ContactName:  c.ContactName,
		ContactTitle: c.ContactTitle,
		Address:      c.Address,
		City:         c.City,
		Region:       c.Region,
		PostalCode:   c.PostalCode,
		Country:      c.Country,
		Phone:        c.Phone,
		Fax:          c.Fax,
	})
	if err != nil {
		os.Exit(1)
	}
}
