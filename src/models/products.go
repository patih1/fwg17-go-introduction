package models

import (
	"fmt"
	"strings"

	"github.com/patih1/fwg17-go-backend/src/services"
)

func FindAllProduct(limit int, offset int, search string) (services.Info, error) {
	// sql := `SELECT "p"."id", "p"."name", "p"."image", "p"."discount", "p"."isRecommended", "p"."basePrice", "p"."description", "p"."createdAt", "p"."updatedAt", "c"."name" as "category" FROM "products" "p"
	// join "productCategories" "pc" on "pc"."productId" = "p"."id"
	// join "categories" "c" on "c"."id" = "pc"."categoryId"
	// WHERE "p"."name" ILIKE $3 ORDER BY "p"."id" ASC LIMIT $1 OFFSET $2`
	sql := `SELECT * FROM "products"
	WHERE "name" ILIKE $3 ORDER BY "id" ASC LIMIT $1 OFFSET $2`
	sqlCount := fmt.Sprintf(`SELECT COUNT(*) FROM "products" where "name" ILIKE '%%%v%%'`, search)
	result := services.Info{}
	data := []services.Product{}
	fmtSearch := fmt.Sprintf("%%%v%%", search)
	err := db.Select(&data, sql, limit, offset, fmtSearch)
	result.Data = data

	fmt.Print(err)

	row := db.QueryRow(sqlCount)
	err = row.Scan(&result.Count)

	return result, err
}

func FindOneProduct(id *int) (services.Product, error) {
	sql := `SELECT * FROM "products" WHERE
	"id" = $1`
	// sql := `SELECT "p"."id", "p"."name", "p"."image", "p"."discount", "p"."isRecommended", "p"."basePrice", "p"."description", "p"."createdAt", "p"."updatedAt", "c"."name" as "category" FROM "products" "p"
	// join "productCategories" "pc" on "pc"."productId" = "p"."id"
	// join "categories" "c" on "c"."id" = "pc"."categoryId"
	// WHERE "p"."id" = $1`
	data := services.Product{}
	err := db.Get(&data, sql, id)
	return data, err
}

func CreateProduct(col []string, values []string) (services.Product, error) {
	sql := fmt.Sprint(`INSERT INTO "products" (`, strings.Join(col, ", "), `) VALUES (`, strings.Join(values, ", "), `) RETURNING *`)
	data := services.Product{}
	err := db.Get(&data, sql)
	return data, err
}

func UpdateProduct(values []string, id int) (services.Product, error) {
	sql := fmt.Sprint(`UPDATE "products" SET`, strings.Join(values, ", "), `, "updatedAt"=now() WHERE "id"=$1 RETURNING *`)
	data := services.Product{}
	err := db.Get(&data, sql, id)
	return data, err
}

func DeleteProduct(id int) (services.Product, error) {
	sql := `DELETE FROM "products" WHERE "id" = $1 RETURNING *`
	data := services.Product{}
	err := db.Get(&data, sql, id)
	return data, err
}
