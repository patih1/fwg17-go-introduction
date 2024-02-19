package models

import (
	"fmt"
	"strings"

	"github.com/patih1/fwg17-go-backend/src/services"
)

func FindAllProductRatings(limit int, offset int) (services.Info, error) {
	sql := `SELECT * FROM "productRatings"
	ORDER BY "id" ASC LIMIT $1 OFFSET $2`
	sqlCount := `SELECT COUNT(*) FROM "productRatings"`
	result := services.Info{}
	data := []services.ProductRatings{}
	err := db.Select(&data, sql, limit, offset)
	result.Data = data

	fmt.Print(err)

	row := db.QueryRow(sqlCount)
	err = row.Scan(&result.Count)

	return result, err
}

func FindOneProductRatings(id *int) (services.ProductRatings, error) {
	sql := `SELECT * FROM "productRatings" WHERE
	"id" = $1`
	data := services.ProductRatings{}
	err := db.Get(&data, sql, id)
	return data, err
}

func CreateProductRatings(col []string, values []string) (services.ProductRatings, error) {
	sql := fmt.Sprint(`INSERT INTO "productRatings" (`, strings.Join(col, ", "), `) VALUES (`, strings.Join(values, ", "), `) RETURNING *`)
	data := services.ProductRatings{}
	err := db.Get(&data, sql)
	return data, err
}

func UpdateProductRatings(values []string, id int) (services.ProductRatings, error) {
	sql := fmt.Sprint(`UPDATE "productRatings" SET`, strings.Join(values, ", "), `, "updatedAt"=now() WHERE "id"=$1 RETURNING *`)
	data := services.ProductRatings{}
	err := db.Get(&data, sql, id)
	return data, err
}

func DeleteProductRatings(id int) (services.ProductRatings, error) {
	sql := `DELETE FROM "productRatings" WHERE "id" = $1 RETURNING *`
	data := services.ProductRatings{}
	err := db.Get(&data, sql, id)
	return data, err
}
