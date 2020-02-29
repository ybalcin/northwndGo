package services

import (
	"context"
	f "github.com/ybalcin/northwndGo/firestore"
	i "github.com/ybalcin/northwndGo/inputs"
	u "github.com/ybalcin/northwndGo/utils"
)

func GetAll(input i.GetInput) ([]map[string]interface{}, error) {
	client := f.Client()
	defer client.Close()
	ctx := context.Background()

	var e error
	var values []map[string]interface{}
	dSnapshot, _ := client.Collection(input.Collection).Documents(ctx).GetAll()
	for _, snapshot := range dSnapshot {
		val, err := u.Map(snapshot, input.Model)
		if err != nil {
			e = err
		}
		values = append(values, val.(map[string]interface{}))
	}
	return values, e
}

func GetById(input i.GetInput) (map[string]interface{}, error) {
	client := f.Client()
	defer client.Close()
	ctx := context.Background()

	var e error
	ds, err := client.Collection(input.Collection).Doc(input.ID).Get(ctx)
	if err != nil {
		e = err
	}
	val, err := u.Map(ds, input.Model)
	if err != nil {
		e = err
	}
	return val.(map[string]interface{}), e
}
