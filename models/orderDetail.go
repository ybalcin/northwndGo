package models

type OrderDetail struct {
	OrderID   *int
	ProductID *int
	UnitPrice *float64
	Quantity  *int16
	Discount  *float64
}
