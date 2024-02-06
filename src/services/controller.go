package services

import (
	"time"

	"github.com/LukaGiorgadze/gonull"
)

type User struct {
	Id          int                        `db:"id" json:"id"`
	FullName    string                     `db:"fullName" json:"fullName" form:"fullName"`
	Email       string                     `db:"email" json:"email" form:"email"`
	Password    string                     `db:"password" json:"password" form:"password"`
	Address     gonull.Nullable[string]    `db:"address" json:"address" form:"address"`
	Picture     gonull.Nullable[string]    `db:"picture" json:"picture" form:"picture"`
	PhoneNumber gonull.Nullable[string]    `db:"phoneNumber" json:"phoneNumber" form:"phoneNumber"`
	Role        int                        `db:"role" json:"role" form:"role"`
	CreatedAt   time.Time                  `db:"createdAt" json:"createdAt"`
	UpdatedAt   gonull.Nullable[time.Time] `db:"updatedAt" json:"updatedAt" form:"updatedAt"`
}

type ToUpdateUser struct {
	Id          int                        `db:"id" json:"id"`
	FullName    string                     `db:"fullName" json:"fullName" form:"fullName"`
	Email       string                     `db:"email" json:"email" form:"email"`
	Password    string                     `db:"password" json:"password" form:"password"`
	Address     string                     `db:"address" json:"address" form:"address"`
	Picture     string                     `db:"picture" json:"picture" form:"picture"`
	PhoneNumber string                     `db:"phoneNumber" json:"phoneNumber" form:"phoneNumber"`
	Role        int                        `db:"role" json:"role" form:"role"`
	CreatedAt   time.Time                  `db:"createdAt" json:"createdAt"`
	UpdatedAt   gonull.Nullable[time.Time] `db:"updatedAt" json:"updatedAt" form:"updatedAt"`
}

type PageInfo struct {
	Page      int `json:"page"`
	Limit     int `json:"limit"`
	LastPage  int `json:"lastPage"`
	TotalData int `json:"totalData"`
}

type Response struct {
	Success  bool        `json:"success"`
	Message  string      `json:"message"`
	Pageinfo PageInfo    `json:"pageInfo"`
	Results  interface{} `json:"results"`
}

type ResponseOnly struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type Login struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Token   string `json:"token"`
}

type Info struct {
	Data  interface{}
	Count int
}
