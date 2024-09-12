package model

type OrderAddress struct {
	OrderID      int     `json:"order_id" db:"order_id"`
	Country      string  `json:"country" db:"country"`
	City         string  `json:"city" db:"city"`
	District     string  `json:"district" db:"district"`
	Ward         *string `json:"ward" db:"ward"`
	Street       *string `json:"street" db:"street"`
	HouseNumber  *string `json:"house_number" db:"house_number"`
	AddressLine1 *string `json:"address_line1" db:"address_line1"`
	AddressLine2 *string `json:"address_line2" db:"address_line2"`
}
