package model

import "time"

type PartnerGroup struct {
	ID        int    `json:"id" db:"id"`
	GroupCode string `json:"group_code" db:"group_code"`
	GroupName string `json:"group_name" db:"group_name"`
	GroupType string `json:"group_type" db:"group_type"`

	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
}

type PartnerGroupPage struct {
	Size       int             `json:"size"`
	Page       int             `json:"page"`
	TotalPage  int             `json:"total_page"`
	TotalCount int             `json:"total_count"`
	Hasmore    bool            `json:"hasmore"`
	Groups     []*PartnerGroup `json:"groups"`
}
