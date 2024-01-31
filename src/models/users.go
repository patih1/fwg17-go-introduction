package models

import (
	"time"

	"github.com/LukaGiorgadze/gonull"
	"github.com/jmoiron/sqlx"
	"github.com/patih1/fwg17-go-backend/src/lib"
)

var db *sqlx.DB = lib.DB

type User struct {
	Id          int                        `db:"id" json:"id"`
	Fullname    string                     `db:"fullName" json:"fullName" form:"fullName"`
	Email       string                     `db:"email" json:"email" form:"email"`
	Password    string                     `db:"password" json:"password" form:"password"`
	Address     gonull.Nullable[string]    `db:"address" json:"address"`
	Picture     gonull.Nullable[string]    `db:"picture" json:"picture"`
	PhoneNumber gonull.Nullable[string]    `db:"phoneNumber" json:"phoneNumber"`
	Role        string                     `db:"role" json:"role"`
	CreatedAt   time.Time                  `db:"createdAt" json:"createdAt"`
	UpdatedAt   gonull.Nullable[time.Time] `db:"updatedAt" json:"updatedAt"`
}

func CreateUser(data User) (User, error) {
	sql := `INSERT INTO "users" ("fullName", "email", "password") VALUES (:fullName, :email, :password) RETURNING *`
	result := User{}
	rows, err := db.NamedQuery(sql, data)

	for rows.Next() {
		rows.StructScan(&result)
	}
	return result, err
}
