package models

type Supplier struct {
	SupplierID   *int
	CompanyName  *string
	ContactName  *string
	ContactTitle *string
	Address      *string
	City         *string
	Region       *string
	PostalCode   *string
	Country      *string
	Phone        *string
	Fax          *string
	HomePage     *string
}
