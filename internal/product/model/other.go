package model

type ProductAttachment struct {
	ProductID   int     `json:"-" db:"product_id"`
	FileID      string  `json:"file_id" db:"file_id"`
	FileName    *string `json:"file_name" db:"file_name"`
	FilePath    *string `json:"-" db:"file_path"`
	FileLink    *string `json:"file_link"`
	Description *string `json:"description" db:"description"`
}

type ProductCategory struct {
	ProductID     int     `json:"-" db:"product_id"`
	CategoryID    int     `json:"category_id" db:"category_id"`
	CategoryName  *string `json:"category_name" db:"category_name"`
	ThumbnailID   *string `json:"thumbnail_id" db:"thumbnail_id"`
	ThumbnailPath *string `json:"-" db:"thumbnail_path"`
	ThumbnailLink *string `json:"thumbnail_link"`
}

type ProductImage struct {
	ProductID   int     `json:"-" db:"product_id"`
	ImageID     *string `json:"image_id" db:"image_id"`
	ImagePath   *string `json:"-" db:"image_path"`
	ImageName   *string `json:"image_name" db:"image_name"`
	ImageLink   *string `json:"image_link"`
	Description *string `json:"description" db:"description"`
}
