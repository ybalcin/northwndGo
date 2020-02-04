package models

type Product struct {
	ProductID       *int
	ProductName     *string
	SupplierID      *int
	CategoryID      *int
	QuantityPerUnit *string
	UnitPrice       *float64
	UnitsInStock    *int16
	UnitsOnOrder    *int16
	ReorderLevel    *int16
	Discontinued    bool
}
