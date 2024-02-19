package models

import (
	"fmt"
	"strings"

	"github.com/patih1/fwg17-go-backend/src/services"
)

func FindAllProductTags(limit int, offset int) (services.Info, error) {
	sql := `SELECT * FROM "productTags"
	ORDER BY "id" ASC LIMIT $1 OFFSET $2`
	sqlCount := `SELECT COUNT(*) FROM "productTags"`
	result := services.Info{}
	data := []services.ProductTags{}
	err := db.Select(&data, sql, limit, offset)
	result.Data = data

	fmt.Print(err)

	row := db.QueryRow(sqlCount)
	err = row.Scan(&result.Count)

	return result, err
}

func FindOneProductTags(id *int) (services.ProductTags, error) {
	sql := `SELECT * FROM "productTags" WHERE
	"id" = $1`
	data := services.ProductTags{}
	err := db.Get(&data, sql, id)
	return data, err
}

func CreateProductTags(col []string, values []string) (services.ProductTags, error) {
	sql := fmt.Sprint(`INSERT INTO "productTags" (`, strings.Join(col, ", "), `) VALUES (`, strings.Join(values, ", "), `) RETURNING *`)
	data := services.ProductTags{}
	err := db.Get(&data, sql)
	return data, err
}

func UpdateProductTags(values []string, id int) (services.ProductTags, error) {
	sql := fmt.Sprint(`UPDATE "productTags" SET`, strings.Join(values, ", "), `, "updatedAt"=now() WHERE "id"=$1 RETURNING *`)
	data := services.ProductTags{}
	err := db.Get(&data, sql, id)
	return data, err
}

func DeleteProductTags(id int) (services.ProductTags, error) {
	sql := `DELETE FROM "productTags" WHERE "id" = $1 RETURNING *`
	data := services.ProductTags{}
	err := db.Get(&data, sql, id)
	return data, err
}
