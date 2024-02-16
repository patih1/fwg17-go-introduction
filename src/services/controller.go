package services

import (
	"time"
)

// Information
type PageInfo struct {
	Page      int `json:"page"`
	Limit     int `json:"limit"`
	LastPage  int `json:"lastPage"`
	TotalData int `json:"totalData"`
}

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	// Pageinfo PageInfo    `json:"pageInfo"`
	Results interface{} `json:"results"`
}

type ResponseWPage struct {
	Success  bool        `json:"success"`
	Message  string      `json:"message"`
	Pageinfo PageInfo    `json:"pageInfo"`
	Results  interface{} `json:"results"`
}

type ResponseOnly struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type Info struct {
	Data  interface{}
	Count int
}

// Auth
type Login struct {
	Id       int     `db:"id" json:"id"`
	Email    *string `db:"email" json:"email" form:"email"`
	Password *string `db:"password" json:"password" form:"password"`
	Role     *string `db:"role" json:"role" form:"role"`
}

// User

type User struct {
	Id          int        `db:"id" json:"id"`
	FullName    *string    `db:"fullName" json:"fullName" form:"fullName"`
	Email       *string    `db:"email" json:"email" form:"email"`
	Password    *string    `db:"password" json:"password" form:"password"`
	Address     *string    `db:"address" json:"address" form:"address"`
	PhoneNumber *string    `db:"phoneNumber" json:"phoneNumber" form:"phoneNumber"`
	Role        *string    `db:"role" json:"role" form:"role"`
	Picture     *string    `db:"picture" json:"picture" form:"picture"`
	CreatedAt   *time.Time `db:"createdAt" json:"createdAt"`
	UpdatedAt   *time.Time `db:"updatedAt" json:"updatedAt" form:"updatedAt"`
}

// Products
type Product struct {
	Id            int        `db:"id" json:"id"`
	Name          *string    `db:"name" json:"name" form:"name"`
	Description   *string    `db:"description" json:"description" form:"description"`
	BasePrice     *int       `db:"basePrice" json:"basePrice" form:"basePrice"`
	Category      *string    `db:"category" json:"category" form:"category"`
	Image         *string    `db:"image" json:"image" form:"image"`
	Discount      *string    `db:"discount" json:"discount" form:"discount"`
	IsRecommended *bool      `db:"isRecommended" json:"isRecommended" form:"isRecommended"`
	CreatedAt     *time.Time `db:"createdAt" json:"createdAt"`
	UpdatedAt     *time.Time `db:"updatedAt" json:"updatedAt"`
}
