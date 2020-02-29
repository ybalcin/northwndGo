package helper

import (
	s "database/sql"
	"encoding/json"
	m "github.com/ybalcin/northwndGo/models"
	"log"
)

func Log(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func ConvertToJson(v interface{}) []byte {
	b, err := json.Marshal(v)
	Log(err)
	return b
}

func RowToCustomer(r *s.Row) (m.Customer, error) {
	var customer m.Customer
	err := r.Scan(&customer.CustomerID, &customer.PostalCode, &customer.Country, &customer.Phone, &customer.Fax,
		&customer.Region, &customer.City, &customer.Address, &customer.ContactTitle, &customer.ContactName,
		&customer.CompanyName)
	return customer, err
}

func RowsToCustomer(r *s.Rows) (m.Customer, error) {
	var customer m.Customer
	err := r.Scan(&customer.CustomerID, &customer.CompanyName, &customer.ContactName,
		&customer.ContactTitle, &customer.Address, &customer.City, &customer.Region, &customer.PostalCode,
		&customer.Country, &customer.Phone, &customer.Fax)
	return customer, err
}

func RowsToEmployee(r *s.Rows) (m.Employee, error) {
	var e m.Employee
	err := r.Scan(&e.EmployeeID, &e.LastName, &e.FirstName, &e.Title, &e.TitleOfCourtesy, &e.BirthDate, &e.HireDate,
		&e.Address, &e.City, &e.Region, &e.PostalCode, &e.Country, &e.HomePhone, &e.Extension, &e.PhotoPath, &e.Notes,
		&e.ReportsTo, &e.PhotoPath)
	return e, err
}

func RowsToCategory(r *s.Rows) (m.Category, error) {
	var c m.Category
	err := r.Scan(&c.CategoryID, &c.CategoryName, &c.Description, &c.Picture)
	return c, err
}

func RowsToEmployeeTerritory(r *s.Rows) (m.EmployeeTerritory, error) {
	var e m.EmployeeTerritory
	err := r.Scan(&e.EmployeeID, &e.TerritoryID)
	return e, err
}

func RowsToOrderDetail(r *s.Rows) (m.OrderDetail, error) {
	var o m.OrderDetail
	err := r.Scan(&o.OrderID, &o.ProductID, &o.UnitPrice, &o.Quantity, &o.Discount)
	return o, err
}

func RowsToOrder(r *s.Rows) (m.Order, error) {
	var o m.Order
	var e error
	if err := r.Scan(&o.OrderID, &o.CustomerID, &o.EmployeeID, &o.OrderDate, &o.RequiredDate, &o.ShippedDate,
		&o.ShipVia, &o.Freight, &o.ShipName, &o.ShipAddress, &o.ShipCity, &o.ShipRegion, &o.ShipPostalCode,
		&o.ShipCountry); err != nil {
		e = err
	}
	return o, e
}

func RowsToProduct(r *s.Rows) (m.Product, error) {
	var p m.Product
	var e error
	if err := r.Scan(&p.ProductID, &p.ProductName, &p.SupplierID, &p.CategoryID, &p.QuantityPerUnit,
		&p.UnitPrice, &p.UnitsInStock, &p.UnitsOnOrder, &p.ReorderLevel, &p.Discontinued); err != nil {
		e = err
	}
	return p, e
}

func RowsToRegion(r *s.Rows) (m.Region, error) {
	var reg m.Region
	var e error
	if err := r.Scan(&reg.RegionID, &reg.RegionDescription); err != nil {
		e = err
	}
	return reg, e
}

func RowsToShipper(r *s.Rows) (m.Shipper, error) {
	var shipper m.Shipper
	var e error
	if err := r.Scan(&shipper.ShipperID, &shipper.CompanyName, &shipper.Phone); err != nil {
		e = err
	}
	return shipper, e
}

func RowsToSupplier(r *s.Rows) (m.Supplier, error) {
	var supplier m.Supplier
	var e error
	if err := r.Scan(&supplier.SupplierID, &supplier.CompanyName, &supplier.ContactName,
		&supplier.ContactTitle, &supplier.Address, &supplier.City, &supplier.Region,
		&supplier.PostalCode, &supplier.Country, &supplier.Phone, &supplier.Fax, &supplier.HomePage); err != nil {
		e = err
	}
	return supplier, e
}

func RowsToTerritory(r *s.Rows) (m.Territory, error) {
	var territory m.Territory
	var e error
	if err := r.Scan(&territory.TerritoryID, &territory.TerritoryDescription, &territory.RegionID); err != nil {
		e = err
	}
	return territory, e
}
