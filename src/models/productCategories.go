package models

import (
	"fmt"
	"strings"

	"github.com/patih1/fwg17-go-backend/src/services"
)

func FindAllProductCategories(limit int, offset int) (services.Info, error) {
	sql := `SELECT * FROM "productCategories"
	ORDER BY "id" ASC LIMIT $1 OFFSET $2`
	sqlCount := `SELECT COUNT(*) FROM "productCategories"`
	result := services.Info{}
	data := []services.ProductCategories{}
	err := db.Select(&data, sql, limit, offset)
	result.Data = data

	fmt.Print(err)

	row := db.QueryRow(sqlCount)
	err = row.Scan(&result.Count)

	return result, err
}

func FindOneProductCategories(id *int) (services.ProductCategories, error) {
	sql := `SELECT * FROM "productCategories" WHERE
	"id" = $1`
	data := services.ProductCategories{}
	err := db.Get(&data, sql, id)
	return data, err
}

func CreateProductCategories(col []string, values []string) (services.ProductCategories, error) {
	sql := fmt.Sprint(`INSERT INTO "productCategories" (`, strings.Join(col, ", "), `) VALUES (`, strings.Join(values, ", "), `) RETURNING *`)
	data := services.ProductCategories{}
	err := db.Get(&data, sql)
	return data, err
}

func UpdateProductCategories(values []string, id int) (services.ProductCategories, error) {
	sql := fmt.Sprint(`UPDATE "productCategories" SET`, strings.Join(values, ", "), `, "updatedAt"=now() WHERE "id"=$1 RETURNING *`)
	data := services.ProductCategories{}
	err := db.Get(&data, sql, id)
	return data, err
}

func DeleteProductCategories(id int) (services.ProductCategories, error) {
	sql := `DELETE FROM "productCategories" WHERE "id" = $1 RETURNING *`
	data := services.ProductCategories{}
	err := db.Get(&data, sql, id)
	return data, err
}
