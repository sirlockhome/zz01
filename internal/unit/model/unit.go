package model

import "time"

type Unit struct {
	ID          int     `json:"id" db:"id"`
	UnitName    string  `json:"unit_name" db:"unit_name"`
	Symbol      string  `json:"symbol" db:"symbol"`
	Description *string `json:"description" db:"description"`
	IsActive    bool    `json:"is_active" db:"is_active"`

	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
}

type UnitPage struct {
	Page       int     `json:"page"`
	Size       int     `json:"size"`
	TotalCount int     `json:"total_count"`
	TotalPage  int     `json:"total_page"`
	Hasmore    bool    `json:"hasmore"`
	Units      []*Unit `json:"units"`
}
