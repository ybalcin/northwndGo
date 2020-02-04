package models

import (
	t "time"
)

type Employee struct {
	EmployeeID      *int
	LastName        *string
	FirstName       *string
	Title           *string
	TitleOfCourtesy *string
	BirthDate       t.Time
	HireDate        t.Time
	Address         *string
	City            *string
	Region          *string
	PostalCode      *string
	Country         *string
	HomePhone       *string
	Extension       *string
	Photo           []byte
	Notes           *string
	ReportsTo       *int
	PhotoPath       *string
}
