package models

import "time"

type Product struct {
	ID         int                  `json:"id" gorm:"primary_key:auto_increment"`
	Name       string               `json:"name" form:"name" gorm:"type: varchar(255)"`
	Buy        string               `json:"buy" gorm:"type:text" form:"buy"`
	Sell       int                  `json:"sell" form:"sell" gorm:"type: int"`
	Image      string               `json:"image" form:"image" gorm:"type: varchar(255)"`
	Qty        int                  `json:"qty" form:"qty"`
	UserID     int                  `json:"user_id" form:"user_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	User       UsersProfileResponse `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Category   []Category           `json:"category" gorm:"many2many:product_categories;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CategoryID []int                `json:"-" form:"category_id" gorm:"-"`
	CreatedAt  time.Time            `json:"-"`
	UpdatedAt  time.Time            `json:"-"`
}

type ProductResponse struct {
	ID         int                  `json:"id"`
	Name       string               `json:"name"`
	Buy        string               `json:"buy"`
	Sell       int                  `json:"sell"`
	Image      string               `json:"image"`
	Qty        int                  `json:"qty"`
	UserID     int                  `json:"-"`
	User       UsersProfileResponse `json:"user"  gorm:"foreignKey:UserID"`
	Category   []Category           `json:"category" gorm:"many2many:product_categories"`
	CategoryID []int                `json:"-" form:"category_id" gorm:"-"`
}

type ProductUserResponse struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Desc   string `json:"desc"`
	Price  int    `json:"price"`
	Image  string `json:"image"`
	Qty    int    `json:"qty"`
	UserID int    `json:"-"`
}

func (ProductResponse) TableName() string {
	return "products"
}

func (ProductUserResponse) TableName() string {
	return "products"
}
