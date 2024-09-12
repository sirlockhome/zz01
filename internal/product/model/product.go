package model

import "time"

type Product struct {
	ID             int        `json:"id" db:"id"`
	Name           string     `json:"name" db:"name"`
	Description    string     `json:"description" db:"description"`
	SKU            string     `json:"sku" db:"sku"`
	Information    string     `json:"information" db:"information"`
	Feature        string     `json:"feature" db:"feature"`
	Specifications string     `json:"specifications" db:"specifications"`
	Price          float64    `json:"price" db:"price"`
	Stock          int        `json:"stock" db:"stock"`
	UnitGroupID    int        `json:"unit_group_id" db:"unit_group_id"`
	UnitGroup      *UnitGroup `json:"unit_group"`

	UnitPriceID   int    `json:"unit_price_id" db:"unit_price_id"`
	UnitPriceName string `json:"unit_price_name" db:"unit_price_name"`

	ThumbnailID   *string `json:"thumbnail_id" db:"thumbnail_id"`
	ThumbnailPath *string `json:"-" db:"thumbnail_path"`
	ThumbnailLink *string `json:"thumbnail_link"`

	ProductCategories  []*ProductCategory   `json:"product_categories,omitempty"`
	ProductImages      []*ProductImage      `json:"product_images,omitempty"`
	ProductAttachments []*ProductAttachment `json:"product_attachments,omitempty"`

	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
}

type ProductRequest struct {
	ID             int     `json:"-" db:"id"`
	Name           string  `json:"name" db:"name"`
	Description    string  `json:"description" db:"description"`
	SKU            string  `json:"sku" db:"sku"`
	Information    string  `json:"information" db:"information"`
	Feature        string  `json:"feature" db:"feature"`
	Specifications string  `json:"specifications" db:"specifications"`
	Price          float64 `json:"price" db:"price"`
	Stock          int     `json:"stock" db:"stock"`
	UnitGroupID    int     `json:"unit_group_id" db:"unit_group_id"`
	ThumbnailID    string  `json:"thumbnail_id" db:"thumbnail_id"`

	ProductCategories  []*ProductCategory   `json:"product_categories"`
	ProductImages      []*ProductImage      `json:"product_images"`
	ProductAttachments []*ProductAttachment `json:"product_attachments"`
}

type ProductPage struct {
	Page       int        `json:"page"`
	Size       int        `json:"size"`
	TotalPage  int        `json:"total_page"`
	TotalCount int        `json:"total_count"`
	Hasmore    bool       `json:"hasmore"`
	Products   []*Product `json:"products"`
}

type UnitGroup struct {
	ID            int     `json:"id" db:"id"`
	UnitGroupName string  `json:"unit_group_name" db:"unit_group_name"`
	BaseUnitID    int     `json:"base_unit_id" db:"base_unit_id"`
	BaseUnitName  string  `json:"base_unit_name" db:"base_unit_name"`
	Symbol        string  `json:"symbol" db:"symbol"`
	Description   *string `json:"description" db:"description"`

	UnitConversions []*UnitConversion `json:"unit_conversions"`
}

type UnitConversion struct {
	ToUnitID   int     `json:"to_unit_id" db:"to_unit_id"`
	ToUnitName *string `json:"to_unit_name" db:"to_unit_name"`
	BaseQty    float64 `json:"base_qty" db:"base_qty"`
	AltQty     float64 `json:"alt_qty" db:"alt_qty"`
}
