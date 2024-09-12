package model

import "time"

type UnitGroup struct {
	ID            int        `json:"id"`
	UnitGroupName string     `json:"unit_group_name"`
	BaseUnitID    int        `json:"base_unit_id"`
	BaseUnitName  string     `json:"base_unit_name"`
	Symbol        string     `json:"symbol"`
	Description   *string    `json:"description"`
	IsActive      bool       `json:"is_active"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     *time.Time `json:"updated_at"`

	UnitConversions []UnitConversion `json:"unit_conversions"` // Chuyển đổi từ interface{} sang slice của UnitConversion
}

// Struct cho phần unit_conversions
type UnitConversion struct {
	ID               int     `json:"id"`
	ToUnitID         int     `json:"to_unit_id"`
	ToUnitName       string  `json:"to_unit_name"`
	ConversionFactor float64 `json:"conversion_factor"`
	Description      *string `json:"description"`
	BaseQty          int     `json:"base_qty"`
	AltQty           int     `json:"alt_qty"`
	CreatedAt        string  `json:"created_at"`
}
