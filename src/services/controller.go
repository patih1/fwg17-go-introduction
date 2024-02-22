package services

import (
	"time"
)

// information
type PageInfo struct {
	Page      int `json:"page"`
	Limit     int `json:"limit"`
	LastPage  int `json:"lastPage"`
	TotalData int `json:"totalData"`
}

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
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

// auth
type Login struct {
	Id       int     `db:"id" json:"id"`
	Email    *string `db:"email" json:"email" form:"email"`
	Password *string `db:"password" json:"password" form:"password"`
	Role     *string `db:"role" json:"role" form:"role"`
}

// user

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

// products
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

type ProductPrice struct {
	BasePrice *int `db:"basePrice" json:"basePrice" form:"basePrice"`
}

// message
type Message struct {
	Id          int        `db:"id" json:"id"`
	RecipientId *int       `db:"recipientId" json:"recipientId" form:"recipientId"`
	SenderId    *int       `db:"senderId" json:"senderId" form:"senderId"`
	Text        *string    `db:"text" json:"text" form:"text"`
	CreatedAt   *time.Time `db:"createdAt" json:"createdAt"`
	UpdatedAt   *time.Time `db:"updatedAt" json:"updatedAt"`
}

type CSMessage struct {
	Id          int     `db:"id" json:"id"`
	RecipientId *int    `db:"recipientId" json:"recipientId" form:"recipientId"`
	SenderId    *int    `db:"senderId" json:"senderId" form:"senderId"`
	Text        *string `db:"text" json:"text" form:"text"`
	FullName    *string `db:"fullName" json:"fullName"`
	Picture     *string `db:"picture" json:"picture" form:"picture"`
}

// orderDetails
type OrderDetails struct {
	Id               int        `db:"id" json:"id"`
	ProductId        *int       `db:"productId" json:"productId" form:"productId"`
	ProductSizeId    *int       `db:"productSizeId" json:"productSizeId" form:"productSizeId"`
	ProductVariantId *int       `db:"productVariantId" json:"productVariantId" form:"productVariantId"`
	Quantity         *int       `db:"quantity" json:"quantity" form:"quantity"`
	OrderId          *int       `db:"orderId" json:"orderId" form:"orderId"`
	SubTotal         *int       `db:"subTotal" json:"subTotal" form:"subTotal"`
	CreatedAt        *time.Time `db:"createdAt" json:"createdAt"`
	UpdatedAt        *time.Time `db:"updatedAt" json:"updatedAt"`
}

type ODWithPDetail struct {
	Id               int     `db:"id" json:"id"`
	Name             *string `db:"name" json:"name"`
	ProductId        *int    `db:"productId" json:"productId"`
	ProductSizeId    *int    `db:"productSizeId" json:"productSizeId"`
	ProductVariantId *int    `db:"productVariantId" json:"productVariantId"`
	Quantity         *int    `db:"quantity" json:"quantity"`
	OrderId          *int    `db:"orderId" json:"orderId"`
	SubTotal         *int    `db:"subTotal" json:"subTotal"`
}

// orders
type Orders struct {
	Id              int        `db:"id" json:"id"`
	UserId          *int       `db:"userId" json:"userId" form:"userId"`
	OrderNumber     *string    `db:"orderNumber" json:"orderNumber" form:"orderNumber"`
	PromoId         *int       `db:"promoId" json:"promoId" form:"promoId"`
	Total           *int       `db:"total" json:"total" form:"total"`
	TaxAmount       *int       `db:"taxAmount" json:"taxAmount" form:"taxAmount"`
	Status          *string    `db:"status" json:"status" form:"status"`
	DeliveryAddress *string    `db:"deliveryAddress" json:"deliveryAddress" form:"deliveryAddress"`
	FullName        *string    `db:"fullName" json:"fullName" form:"fullName"`
	Email           *string    `db:"email" json:"email" form:"email"`
	CreatedAt       *time.Time `db:"createdAt" json:"createdAt"`
	UpdatedAt       *time.Time `db:"updatedAt" json:"updatedAt"`
}

// productCategories
type ProductCategories struct {
	Id         int        `db:"id" json:"id"`
	ProductId  *int       `db:"productId" json:"productId" form:"productId"`
	CategoryId *int       `db:"categoryId" json:"categoryId" form:"categoryId"`
	CreatedAt  *time.Time `db:"createdAt" json:"createdAt"`
	UpdatedAt  *time.Time `db:"updatedAt" json:"updatedAt"`
}

// productRatings
type ProductRatings struct {
	Id            int        `db:"id" json:"id"`
	ProductId     *int       `db:"productId" json:"productId" form:"productId"`
	Rate          *int       `db:"rate" json:"rate" form:"rate"`
	ReviewMessage *string    `db:"reviewMessage" json:"reviewMessage" form:"reviewMessage"`
	UserId        *int       `db:"userId" json:"userId" form:"userId"`
	CreatedAt     *time.Time `db:"createdAt" json:"createdAt"`
	UpdatedAt     *time.Time `db:"updatedAt" json:"updatedAt"`
}

// productSize
type ProductSize struct {
	Id              int        `db:"id" json:"id"`
	Size            *string    `db:"size" json:"size" form:"size"`
	AdditionalPrice *int       `db:"additionalPrice" json:"additionalPrice" form:"additionalPrice"`
	CreatedAt       *time.Time `db:"createdAt" json:"createdAt"`
	UpdatedAt       *time.Time `db:"updatedAt" json:"updatedAt"`
}
type PSAdditionalPrice struct {
	AdditionalPrice *int `db:"additionalPrice" json:"additionalPrice" form:"additionalPrice"`
}

// productTags
type ProductTags struct {
	Id        int        `db:"id" json:"id"`
	TagId     *int       `db:"tagId" json:"tagId" form:"tagId"`
	ProductId *int       `db:"productId" json:"productId" form:"productId"`
	CreatedAt *time.Time `db:"createdAt" json:"createdAt"`
	UpdatedAt *time.Time `db:"updatedAt" json:"updatedAt"`
}

// productVariant
type ProductVariant struct {
	Id              int        `db:"id" json:"id"`
	Name            *string    `db:"name" json:"name" form:"name"`
	AdditionalPrice *int       `db:"additionalPrice" json:"additionalPrice" form:"additionalPrice"`
	CreatedAt       *time.Time `db:"createdAt" json:"createdAt"`
	UpdatedAt       *time.Time `db:"updatedAt" json:"updatedAt"`
}
type PVAdditionalPrice struct {
	AdditionalPrice *int `db:"additionalPrice" json:"additionalPrice" form:"additionalPrice"`
}

// promo
type Promo struct {
	Id            int        `db:"id" json:"id"`
	Name          *string    `db:"name" json:"name" form:"name"`
	Code          *string    `db:"code" json:"code" form:"code"`
	Description   *string    `db:"description" json:"description" form:"description"`
	Precentage    *float64   `db:"percentage" json:"percentage" form:"percentage"`
	MaximumPromo  *int       `db:"maximumPromo" json:"maximumPromo" form:"maximumPromo"`
	MinimumAmount *int       `db:"minimumAmount" json:"minimumAmount" form:"minimumAmount"`
	IsExpired     *bool      `db:"isExpired" json:"isExpired" form:"isExpired"`
	CreatedAt     *time.Time `db:"createdAt" json:"createdAt"`
	UpdatedAt     *time.Time `db:"updatedAt" json:"updatedAt"`
}

type Tags struct {
	Id        int        `db:"id" json:"id"`
	Name      *string    `db:"name" json:"name" form:"name"`
	CreatedAt *time.Time `db:"createdAt" json:"createdAt"`
	UpdatedAt *time.Time `db:"updatedAt" json:"updatedAt"`
}
