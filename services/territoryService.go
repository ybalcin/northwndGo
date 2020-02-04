package services

import (
	"golang.org/x/net/context"
	"northwndGo/db"
	f "northwndGo/firestore"
	h "northwndGo/helper"
	i "northwndGo/inputs"
	m "northwndGo/models"
)

func GetAllTerritoriesFromDb() ([]m.Territory, error) {
	var e error
	var territories []m.Territory
	db, _ := db.ConnectDb()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM Territories")
	if err != nil {
		e = err
	}
	for rows.Next() {
		t, err := h.RowsToTerritory(rows)
		if err != nil {
			e = err
		}
		territories = append(territories, t)
	}
	return territories, e
}

func GetAllTerritories() ([]map[string]interface{}, error) {
	var e error
	territories, err := GetAll(i.GetInput{
		Collection: "territories",
		Model:      m.Territory{},
	})
	if err != nil {
		e = err
	}
	return territories, e
}

func GetTerritory(id string) (map[string]interface{}, error) {
	var e error
	territory, err := GetById(i.GetInput{
		Collection: "territories",
		Model:      m.Territory{},
		ID:         id,
	})
	if err != nil {
		e = err
	}
	return territory, e
}

func WriteToFirestoreTerritory(t m.Territory) {
	ctx := context.Background()
	client := f.Client()
	defer client.Close()
	_, _, err := client.Collection("territories").Add(ctx, m.Territory{
		TerritoryID:          t.TerritoryID,
		TerritoryDescription: t.TerritoryDescription,
		RegionID:             t.RegionID,
	})
	h.Log(err)
}
