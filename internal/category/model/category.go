package model

import (
	"time"
)

type Category struct {
	ID           int     `json:"id" db:"id"`
	CategoryName string  `json:"category_name" db:"category_name"`
	Description  *string `json:"description" db:"description"`
	ParentID     int     `json:"parent_id" db:"parent_id"`
	ThumbnailID  string  `json:"thumbnail_id" db:"thumbnail_id"`

	Children []*Category `json:"children" db:"children"`

	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
}

type CategoryPage struct {
	Page       int         `json:"page"`
	Size       int         `json:"size"`
	TotalPage  int         `json:"total_page"`
	TotalCount int         `json:"total_count"`
	Hasmore    bool        `json:"hasmore"`
	Categories []*Category `json:"categories"`
}
