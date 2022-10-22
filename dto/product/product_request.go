package productdto

type ProductRequest struct {
	Name       string `json:"name" form:"name" gorm:"type: varchar(255)" validate:"required"`
	Buy        string `json:"buy" gorm:"type:text" form:"buy" validate:"required"`
	Sell       int    `json:"sell" form:"sell" gorm:"type: int" validate:"required"`
	Qty        int    `json:"qty" form:"qty" gorm:"type: int" validate:"required"`
	UserID     int    `json:"user_id" form:"user_id"`
	CategoryID []int  `json:"category_id" form:"category_id" gorm:"type: int"`
}
