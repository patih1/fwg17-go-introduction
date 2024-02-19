package models

import (
	"fmt"
	"strings"

	"github.com/patih1/fwg17-go-backend/src/services"
)

func FindAllProductSize(limit int, offset int) (services.Info, error) {
	sql := `SELECT * FROM "productSize"
	ORDER BY "id" ASC LIMIT $1 OFFSET $2`
	sqlCount := `SELECT COUNT(*) FROM "productSize"`
	result := services.Info{}
	data := []services.ProductSize{}
	err := db.Select(&data, sql, limit, offset)
	result.Data = data

	fmt.Print(err)

	row := db.QueryRow(sqlCount)
	err = row.Scan(&result.Count)

	return result, err
}

func FindOneProductSize(id *int) (services.ProductSize, error) {
	sql := `SELECT * FROM "productSize" WHERE
	"id" = $1`
	data := services.ProductSize{}
	err := db.Get(&data, sql, id)
	return data, err
}

func FindPSAdditionalPrice(id *int) (services.PSAdditionalPrice, error) {
	sql := `SELECT "additionalPrice" FROM "productSize" WHERE
	"id" = $1`
	data := services.PSAdditionalPrice{}
	err := db.Get(&data, sql, id)
	return data, err
}

func CreateProductSize(col []string, values []string) (services.ProductSize, error) {
	sql := fmt.Sprint(`INSERT INTO "productSize" (`, strings.Join(col, ", "), `) VALUES (`, strings.Join(values, ", "), `) RETURNING *`)
	data := services.ProductSize{}
	err := db.Get(&data, sql)
	return data, err
}

func UpdateProductSize(values []string, id int) (services.ProductSize, error) {
	sql := fmt.Sprint(`UPDATE "productSize" SET`, strings.Join(values, ", "), `, "updatedAt"=now() WHERE "id"=$1 RETURNING *`)
	data := services.ProductSize{}
	err := db.Get(&data, sql, id)
	return data, err
}

func DeleteProductSize(id int) (services.ProductSize, error) {
	sql := `DELETE FROM "productSize" WHERE "id" = $1 RETURNING *`
	data := services.ProductSize{}
	err := db.Get(&data, sql, id)
	return data, err
}
