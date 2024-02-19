package models

import (
	"fmt"
	"strings"

	"github.com/patih1/fwg17-go-backend/src/services"
)

func FindAllProductVariant(limit int, offset int) (services.Info, error) {
	sql := `SELECT * FROM "productVariant"
	ORDER BY "id" ASC LIMIT $1 OFFSET $2`
	sqlCount := `SELECT COUNT(*) FROM "productVariant"`
	result := services.Info{}
	data := []services.ProductVariant{}
	err := db.Select(&data, sql, limit, offset)
	result.Data = data

	fmt.Print(err)

	row := db.QueryRow(sqlCount)
	err = row.Scan(&result.Count)

	return result, err
}

func FindOneProductVariant(id *int) (services.ProductVariant, error) {
	sql := `SELECT * FROM "productVariant" WHERE
	"id" = $1`
	data := services.ProductVariant{}
	err := db.Get(&data, sql, id)
	return data, err
}

func FindPVAdditionalPrice(id *int) (services.PVAdditionalPrice, error) {
	sql := `SELECT "additionalPrice" FROM "productVariant" WHERE
	"id" = $1`
	data := services.PVAdditionalPrice{}
	err := db.Get(&data, sql, id)
	return data, err
}

func CreateProductVariant(col []string, values []string) (services.ProductVariant, error) {
	sql := fmt.Sprint(`INSERT INTO "productVariant" (`, strings.Join(col, ", "), `) VALUES (`, strings.Join(values, ", "), `) RETURNING *`)
	data := services.ProductVariant{}
	err := db.Get(&data, sql)
	return data, err
}

func UpdateProductVariant(values []string, id int) (services.ProductVariant, error) {
	sql := fmt.Sprint(`UPDATE "productVariant" SET`, strings.Join(values, ", "), `, "updatedAt"=now() WHERE "id"=$1 RETURNING *`)
	data := services.ProductVariant{}
	err := db.Get(&data, sql, id)
	return data, err
}

func DeleteProductVariant(id int) (services.ProductVariant, error) {
	sql := `DELETE FROM "productVariant" WHERE "id" = $1 RETURNING *`
	data := services.ProductVariant{}
	err := db.Get(&data, sql, id)
	return data, err
}
