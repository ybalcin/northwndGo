package services

import (
	"context"
	"northwndGo/db"
	f "northwndGo/firestore"
	"northwndGo/helper"
	i "northwndGo/inputs"
	m "northwndGo/models"
	"os"
)

func GetAllEmployeesFromDb() ([]m.Employee, error) {
	var employees []m.Employee
	var e error
	Db, _ := db.ConnectDb()
	defer Db.Close()

	rows, err := Db.Query("SELECT * FROM Employees")
	if err != nil {
		e = err
	}
	for rows.Next() {
		employee, err := helper.RowsToEmployee(rows)
		if err != nil {
			e = err
		}
		employees = append(employees, employee)
	}

	return employees, e
}

func GetAllEmployees() ([]map[string]interface{}, error) {
	var e error
	employees, err := GetAll(i.GetInput{
		Collection: "employees",
		Model:      m.Employee{},
	})
	if err != nil {
		e = err
	}
	return employees, e
}

func GetEmployee(id string) (map[string]interface{}, error) {
	var e error
	employee, err := GetById(i.GetInput{
		Collection: "employees",
		Model:      m.Employee{},
		ID:         id,
	})
	if err != nil {
		e = err
	}
	return employee, e
}

func WriteToFirestoreEmployee(e m.Employee) {
	ctx := context.Background()
	client := f.Client()
	defer client.Close()

	_, _, err := client.Collection("employees").Add(ctx, m.Employee{
		EmployeeID:      e.EmployeeID,
		LastName:        e.LastName,
		FirstName:       e.FirstName,
		Title:           e.Title,
		TitleOfCourtesy: e.TitleOfCourtesy,
		BirthDate:       e.BirthDate,
		HireDate:        e.HireDate,
		Address:         e.Address,
		City:            e.City,
		Region:          e.Region,
		PostalCode:      e.PostalCode,
		Country:         e.Country,
		HomePhone:       e.HomePhone,
		Extension:       e.Extension,
		Photo:           e.Photo,
		Notes:           e.Notes,
		ReportsTo:       e.ReportsTo,
		PhotoPath:       e.PhotoPath,
	})
	if err != nil {
		os.Exit(1)
	}
}
