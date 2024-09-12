package model

import "time"

type UnitGroup struct {
	ID            int        `json:"id" db:"id"`
	UnitGroupName string     `json:"unit_group_name" db:"unit_group_name"`
	BaseUnitID    int        `json:"base_unit_id" db:"base_unit_id"`
	BaseUnitName  string     `json:"base_unit_name" db:"base_unit_name"`
	Symbol        string     `json:"symbol" db:"symbol"`
	Description   *string    `json:"description" db:"description"`
	IsActive      bool       `json:"is_active" db:"is_active"`
	CreatedAt     time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt     *time.Time `json:"updated_at" db:"updated_at"`

	UnitConversions []*UnitConversion `json:"unit_conversions"`
}

type UnitGroupPage struct {
	Page       int          `json:"page"`
	Size       int          `json:"size"`
	TotalPage  int          `json:"total_page"`
	TotalCount int          `json:"total_count"`
	Hasmore    bool         `json:"hasmore"`
	UnitGroups []*UnitGroup `json:"unit_groups"`
}

type UnitConversion struct {
	ID               int     `json:"id" db:"id"`
	UnitGroupID      int     `json:"unit_group_id" db:"unit_group_id"`
	ToUnitID         int     `json:"to_unit_id" db:"to_unit_id"`
	ToUnitName       *string `json:"to_unit_name" db:"to_unit_name"`
	ConversionFactor float64 `json:"conversion_factor" db:"conversion_factor"`
	BaseQty          float64 `json:"base_qty" db:"base_qty"`
	AltQty           float64 `json:"alt_qty" db:"alt_qty"`
}
