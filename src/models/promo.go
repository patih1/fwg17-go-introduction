package models

import (
	"fmt"
	"strings"

	"github.com/patih1/fwg17-go-backend/src/services"
)

func FindAllPromo(limit int, offset int) (services.Info, error) {
	sql := `SELECT * FROM "promo"
	ORDER BY "id" ASC LIMIT $1 OFFSET $2`
	sqlCount := `SELECT COUNT(*) FROM "promo"`
	result := services.Info{}
	data := []services.Promo{}
	err := db.Select(&data, sql, limit, offset)
	result.Data = data

	fmt.Print(err)

	row := db.QueryRow(sqlCount)
	err = row.Scan(&result.Count)

	return result, err
}

func FindOnePromo(id *int) (services.Promo, error) {
	sql := `SELECT * FROM "promo" WHERE
	"id" = $1`
	data := services.Promo{}
	err := db.Get(&data, sql, id)
	return data, err
}

func CreatePromo(col []string, values []string) (services.Promo, error) {
	sql := fmt.Sprint(`INSERT INTO "promo" (`, strings.Join(col, ", "), `) VALUES (`, strings.Join(values, ", "), `) RETURNING *`)
	data := services.Promo{}
	err := db.Get(&data, sql)
	return data, err
}

func UpdatePromo(values []string, id int) (services.Promo, error) {
	sql := fmt.Sprint(`UPDATE "promo" SET`, strings.Join(values, ", "), `, "updatedAt"=now() WHERE "id"=$1 RETURNING *`)
	data := services.Promo{}
	err := db.Get(&data, sql, id)
	return data, err
}

func DeletePromo(id int) (services.Promo, error) {
	sql := `DELETE FROM "promo" WHERE "id" = $1 RETURNING *`
	data := services.Promo{}
	err := db.Get(&data, sql, id)
	return data, err
}
