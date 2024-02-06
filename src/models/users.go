package models

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/patih1/fwg17-go-backend/src/lib"
	"github.com/patih1/fwg17-go-backend/src/services"
)

var db *sqlx.DB = lib.DB

func FindAll(limit int, offset int, search string) (services.Info, error) {
	sql := `SELECT * FROM "users" WHERE "fullName" ILIKE $3 ORDER BY "id" ASC LIMIT $1 OFFSET $2`
	sqlCount := `SELECT COUNT(*) FROM "users"`
	result := services.Info{}
	data := []services.User{}
	fmtSearch := fmt.Sprintf("%%%v%%", search)
	err := db.Select(&data, sql, limit, offset, fmtSearch)
	result.Data = data

	row := db.QueryRow(sqlCount)
	err = row.Scan(&result.Count)
	// result.Count = 0

	return result, err
}

func FindOne(id int) (services.User, error) {
	sql := `SELECT * FROM "users" WHERE "id" = $1`
	data := services.User{}
	err := db.Get(&data, sql, id)
	return data, err
}

func FindEmail(email string) (services.User, error) {
	sql := `SELECT * FROM "users" WHERE "email" = $1`
	data := services.User{}
	err := db.Get(&data, sql, email)
	return data, err
}

func Create(data services.User) (services.User, error) {
	sql := `INSERT INTO "users" ("fullName", "email", "password", "address") VALUES (:fullName, :email, :password, :address
		) RETURNING *`
	result := services.User{}
	// fmt.Println(result)
	rows, err := db.NamedQuery(sql, data)

	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}
	fmt.Println(rows)
	return result, err
}

// "role"=COALESCE(NULLIF(:role,''),"role"),

func Update(data services.ToUpdateUser) (services.ToUpdateUser, error) {
	sql := `UPDATE "users" SET 
	"fullName"=COALESCE(NULLIF(:fullName,''),"fullName"),
	"email"=COALESCE(NULLIF(:email,''),"email"),
	"password"=COALESCE(NULLIF(:password,''),"password"),
	"address"=COALESCE(NULLIF(:address,''),"address"),
	"picture"=COALESCE(NULLIF(:picture,''),"picture"),
	"phoneNumber"=COALESCE(NULLIF(:phoneNumber,''),"phoneNumber"),
	"updatedAt"=now()
	WHERE "id"=:id
	RETURNING *
	`

	result := services.ToUpdateUser{}
	rows, err := db.NamedQuery(sql, data)

	for rows.Next() {
		rows.StructScan(&result)
	}
	return result, err
}

func Delete(id int) (services.User, error) {
	sql := `DELETE FROM "users" WHERE "id" = $1 RETURNING *`
	data := services.User{}
	err := db.Get(&data, sql, id)
	return data, err
}

func DynamicCreate(col []string, values []string) (services.User, error) {

	// var col []string
	sql := fmt.Sprint(`INSERT INTO "users" (`, strings.Join(col, `",`), `") VALUES (`, strings.Join(values, ", "), `) RETURNING *`)
	data := services.User{}
	err := db.Get(&data, sql)
	return data, err
	// fmt.Println(result)
	// rows, err := db.NamedQuery(sql, data)

	// for rows.Next() {
	// 	rows.StructScan(&result)
	// }
	// fmt.Println(rows)
	// return result, err

}
