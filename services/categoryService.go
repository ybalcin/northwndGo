package services

import (
	"context"
	"github.com/ybalcin/northwndGo/db"
	f "github.com/ybalcin/northwndGo/firestore"
	"github.com/ybalcin/northwndGo/helper"
	i "github.com/ybalcin/northwndGo/inputs"
	m "github.com/ybalcin/northwndGo/models"
)

func GetAllCategoriesFromDb() ([]m.Category, error) {
	var categories []m.Category
	var e error
	db, _ := db.ConnectDb()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM Categories")
	if err != nil {
		e = err
	}
	for rows.Next() {
		c, err := helper.RowsToCategory(rows)
		if err != nil {
			e = err
		}
		categories = append(categories, c)
	}
	return categories, e
}

func GetAllCategories() ([]map[string]interface{}, error) {
	var e error
	categories, err := GetAll(i.GetInput{
		Collection: "categories",
		Model:      m.Category{},
	})
	if err != nil {
		e = err
	}
	return categories, e
}

func GetCategory(id string) (map[string]interface{}, error) {
	var e error
	category, err := GetById(i.GetInput{
		Collection: "categories",
		Model:      m.Category{},
		ID:         id,
	})
	if err != nil {
		e = err
	}
	return category, e
}

func WriteToFirebaseCategory(c m.Category) {
	client := f.Client()
	ctx := context.Background()
	_, _, err := client.Collection("categories").Add(ctx, m.Category{
		CategoryID:   c.CategoryID,
		CategoryName: c.CategoryName,
		Description:  c.Description,
		Picture:      c.Picture,
	})
	helper.Log(err)
}
