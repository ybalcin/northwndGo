package services

import (
	"context"
	"northwndGo/db"
	f "northwndGo/firestore"
	h "northwndGo/helper"
	i "northwndGo/inputs"
	m "northwndGo/models"
)

func GetAllEmployeeTerritoriesFromDb() ([]m.EmployeeTerritory, error) {
	db, _ := db.ConnectDb()
	var e error
	var employeeTerritories []m.EmployeeTerritory
	defer db.Close()

	rows, err = db.Query("SELECT	* FROM EmployeeTerritories")
	if err != nil {
		e = err
	}
	for rows.Next() {
		eTerritory, err := h.RowsToEmployeeTerritory(rows)
		if err != nil {
			e = err
		}
		employeeTerritories = append(employeeTerritories, eTerritory)
	}
	return employeeTerritories, e
}

func GetAllEmployeeTerritories() ([]map[string]interface{}, error) {
	var e error
	employeeTerritories, err := GetAll(i.GetInput{
		Collection: "employeeTerritories",
		Model:      m.EmployeeTerritory{},
	})
	if err != nil {
		e = err
	}
	return employeeTerritories, e
}

func GetEmployeeTerritory(id string) (map[string]interface{}, error) {
	var e error
	employeeTer, err := GetById(i.GetInput{
		Collection: "employeeTerritories",
		Model:      m.EmployeeTerritory{},
		ID:         id,
	})
	if err != nil {
		e = err
	}
	return employeeTer, e
}

func WriteToFirestoreEmpTerritories(et m.EmployeeTerritory) {
	client := f.Client()
	ctx := context.Background()
	defer client.Close()

	_, _, err := client.Collection("employeeTerritories").Add(ctx, m.EmployeeTerritory{
		EmployeeID:  et.EmployeeID,
		TerritoryID: et.TerritoryID,
	})
	h.Log(err)
}
