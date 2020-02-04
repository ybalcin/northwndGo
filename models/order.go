package models

import t "time"

type Order struct {
	OrderID        *int
	CustomerID     *string
	EmployeeID     *int
	OrderDate      *t.Time
	RequiredDate   *t.Time
	ShippedDate    *t.Time
	ShipVia        *int
	Freight        *float64
	ShipName       *string
	ShipAddress    *string
	ShipCity       *string
	ShipRegion     *string
	ShipPostalCode *string
	ShipCountry    *string
}
