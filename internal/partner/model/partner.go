package model

import "time"

type Partner struct {
	ID          int     `json:"id" db:"id"`
	PartnerName string  `json:"partner_name" db:"partner_name"`
	PartnerCode string  `json:"partner_code" db:"partner_code"`
	PartnerType string  `json:"partner_type" db:"partner_type"`
	AvatarID    *string `json:"avatar_id" db:"avatar_id"`
	TaxID       *string `json:"tax_id" db:"tax_id"`
	Email       *string `json:"email" db:"email"`
	PhoneNumber *string `json:"phone_number" db:"phone_number"`

	Country     *string `json:"country" db:"country"`
	City        *string `json:"city" db:"city"`
	District    *string `json:"district" db:"district"`
	Ward        *string `json:"ward" db:"ward"`
	Street      *string `json:"street" db:"street"`
	HouseNumber *string `json:"house_number" db:"house_number"`

	IsActive bool `json:"is_active" db:"is_active"`

	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
}

type PartnerRequest struct {
	ID          int    `json:"-" db:"id"`
	PartnerName string `json:"partner_name" db:"partner_name"`
	PartnerCode string `json:"partner_code" db:"partner_code"`
	PartnerType string `json:"partner_type" db:"partner_type"`
	AvatarID    string `json:"avatar_id" db:"avatar_id"`
	TaxID       string `json:"tax_id" db:"tax_id"`
	Email       string `json:"email" db:"email"`
	PhoneNumber string `json:"phone_number" db:"phone_number"`

	Country     string `json:"country" db:"country"`
	City        string `json:"city" db:"city"`
	District    string `json:"district" db:"district"`
	Ward        string `json:"ward" db:"ward"`
	Street      string `json:"street" db:"street"`
	HouseNumber string `json:"house_number" db:"house_number"`
}

type PartnerPage struct {
	Page       int        `json:"page"`
	Size       int        `json:"size"`
	TotalCount int        `json:"total_count"`
	TotalPage  int        `json:"total_page"`
	Hasmore    bool       `json:"hasmore"`
	Partners   []*Partner `json:"partners"`
}
