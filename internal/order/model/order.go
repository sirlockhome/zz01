package model

import "time"

type Order struct {
	ID               int        `json:"id" db:"id"`
	UserID           int        `json:"user_id" db:"user_id"`
	Status           string     `json:"status" db:"status"`
	OrderDate        *time.Time `json:"order_date" db:"order_date"`
	BuyerNote        *string    `json:"buyer_note" db:"buyer_note"`
	OrderTrackingID  string     `json:"order_tracking_id" db:"order_tracking_id"`
	TotalAmount      float64    `json:"total_amount"`
	PartnerAddressID int        `json:"partner_address_id"`

	ShippingAddress *OrderAddress  `json:"shipping_address,omitempty"`
	BusinessBuyer   *BusinessBuyer `json:"business_buyer,omitempty"`
	OrderItems      []*OrderItem   `json:"order_items,omitempty"`
}

type OrderItem struct {
	ID            int     `json:"id" db:"id"`
	OrderID       int     `json:"order_id" db:"order_id"`
	ProductID     int     `json:"product_id" db:"product_id"`
	ProductName   string  `json:"product_name" db:"product_name"`
	ProductSKU    string  `json:"product_sku" db:"product_sku"`
	ThumbnailID   string  `json:"product_thumbnail_id" db:"thumbnail_id"`
	ThumbnailPath *string `json:"-" db:"thumbnail_path"`
	ThumbnailLink *string `json:"product_thumbnail_link"`
	Quantity      int     `json:"quantity" db:"quantity"`
	Price         float64 `json:"price" db:"price"`
	UnitID        int     `json:"unit_id" db:"unit_id"`
	UnitName      string  `json:"unit_name" db:"unit_name"`

	TotalAmount float64 `json:"total_amount"`
}

type OrderPage struct {
	Page      int      `json:"page"`
	Size      int      `json:"size"`
	TotalPage int      `json:"total_page"`
	TotalSize int      `json:"total_size"`
	Hasmore   bool     `json:"hasmore"`
	Orders    []*Order `json:"orders"`
}

type BusinessBuyer struct {
	BusinessName    string `json:"business_name" db:"business_name"`
	TaxID           string `json:"tax_id" db:"tax_id"`
	ContactName     string `json:"contact_name" db:"contact_name"`
	ContactPhone    string `json:"contact_phone" db:"contact_phone"`
	ContactPosition string `json:"contact_position" db:"contact_position"`
}
