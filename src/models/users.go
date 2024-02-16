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
	sqlCount := fmt.Sprintf(`SELECT COUNT(*) FROM "users" WHERE "fullName" ILIKE '%%%v%%'`, search)
	result := services.Info{}
	data := []services.User{}
	fmtSearch := fmt.Sprintf("%%%v%%", search)
	err := db.Select(&data, sql, limit, offset, fmtSearch)
	result.Data = data

	fmt.Print(err)

	row := db.QueryRow(sqlCount)
	err = row.Scan(&result.Count)
	// result.Count = 0

	return result, err
}

func FindOne(id *int) (services.User, error) {
	sql := `SELECT * FROM "users" WHERE "id" = $1`
	data := services.User{}
	err := db.Get(&data, sql, id)
	return data, err
}

func FindEmail(email *string) (services.User, error) {
	sql := `SELECT * FROM "users" WHERE "email" = $1`
	data := services.User{}
	err := db.Get(&data, sql, email)
	return data, err
}

func Create(data services.User) (services.User, error) {
	sql := `
	INSERT INTO "users" ("fullName", "email", "password") 
	VALUES 
	(:fullName, :email, :password) 
	RETURNING *`
	result := services.User{}
	rows, err := db.NamedQuery(sql, data)

	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}
	return result, err
}

func DynamicCreate(col []string, values []string) (services.User, error) {
	sql := fmt.Sprint(`INSERT INTO "users" (`, strings.Join(col, ", "), `) VALUES (`, strings.Join(values, ", "), `) RETURNING *`)
	data := services.User{}
	err := db.Get(&data, sql)
	return data, err
}

func UpdatePassword(data services.User) (services.User, error) {
	sql := `UPDATE "users" SET 
	"password"=COALESCE(NULLIF(:password,''),"password"),
	"updatedAt"=now()
	WHERE "id"=:id
	RETURNING *
	`

	result := services.User{}
	rows, err := db.NamedQuery(sql, data)

	for rows.Next() {
		rows.StructScan(&result)
	}
	return result, err
}

func DynamicUpdate(values []string, id int) (services.User, error) {
	sql := fmt.Sprint(`UPDATE "users" SET`, strings.Join(values, ", "), `, "updatedAt"=now() WHERE "id"=$1 RETURNING *`)
	data := services.User{}
	err := db.Get(&data, sql, id)
	return data, err
}

func Delete(id int) (services.User, error) {
	sql := `DELETE FROM "users" WHERE "id" = $1 RETURNING *`
	data := services.User{}
	err := db.Get(&data, sql, id)
	return data, err
}
