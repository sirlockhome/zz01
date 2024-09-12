package model

type Address struct {
	ID        int `json:"id" db:"id"`
	PartnerID int `json:"partner_id" db:"partner_id"`

	Country     string  `json:"country" db:"country"`
	City        string  `json:"city" db:"city"`
	District    string  `json:"district" db:"district"`
	Ward        *string `json:"ward" db:"ward"`
	Street      *string `json:"street" db:"street"`
	HouseNumber *string `json:"house_number" db:"house_number"`

	IsDefaultShippingAddress bool `json:"is_default_shipping_address" db:"is_default_shipping_address"`
	IsBillingAddress         bool `json:"is_billing_address" db:"is_billing_address"`

	AddressLine1 string  `json:"address_line1" db:"address_line1"`
	AddressLine2 *string `json:"address_line2" db:"address_line2"`
}
