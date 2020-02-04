package services

import (
	"context"
	"northwndGo/db"
	f "northwndGo/firestore"
	h "northwndGo/helper"
	i "northwndGo/inputs"
	m "northwndGo/models"
	"strings"
)

func GetAllRegionFromDb() ([]m.Region, error) {
	var e error
	var regions []m.Region
	db, _ := db.ConnectDb()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM Region")
	if err != nil {
		e = err
	}
	for rows.Next() {
		reg, err := h.RowsToRegion(rows)
		if err != nil {
			e = err
		}
		regions = append(regions, reg)
	}
	return regions, e
}

func GetAllRegions() ([]map[string]interface{}, error) {
	var e error
	regions, err := GetAll(i.GetInput{
		Collection: "regions",
		Model:      m.Region{},
	})
	if err != nil {
		e = err
	}
	return regions, e
}

func GetRegion(id string) (map[string]interface{}, error) {
	var e error
	region, err := GetById(i.GetInput{
		Collection: "regions",
		Model:      m.Region{},
		ID:         id,
	})
	if err != nil {
		e = err
	}
	return region, e
}

func WriteToFirestoreRegion(r m.Region) {
	ctx := context.Background()
	client := f.Client()
	defer client.Close()
	_, _, err := client.Collection("regions").Add(ctx, m.Region{
		RegionID:          r.RegionID,
		RegionDescription: strings.TrimSpace(r.RegionDescription),
	})
	h.Log(err)
}
