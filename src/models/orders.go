package models

import (
	"fmt"
	"strings"

	"github.com/patih1/fwg17-go-backend/src/services"
)

func FindAllOrders(limit int, offset int) (services.Info, error) {
	sql := `SELECT * FROM "orders"
	ORDER BY "id" ASC LIMIT $1 OFFSET $2`
	sqlCount := `SELECT COUNT(*) FROM "orders"`
	result := services.Info{}
	data := []services.Orders{}
	err := db.Select(&data, sql, limit, offset)
	result.Data = data

	fmt.Print(err)

	row := db.QueryRow(sqlCount)
	err = row.Scan(&result.Count)

	return result, err
}

func FindOneOrders(id *int) (services.Orders, error) {
	sql := `SELECT * FROM "orders" WHERE
	"id" = $1`
	data := services.Orders{}
	err := db.Get(&data, sql, id)
	return data, err
}

func FindHistoryOrder(limit int, offset int, id int) (services.Info, error) {
	sql := `SELECT * FROM "orders" where "userId" = $3
	ORDER BY "id" ASC LIMIT $1 OFFSET $2`
	sqlCount := fmt.Sprintf(`SELECT COUNT(*) FROM "orders" WHERE "userId" = %v`, id)
	result := services.Info{}
	data := []services.Orders{}
	err := db.Select(&data, sql, limit, offset, id)
	result.Data = data

	fmt.Print(err)

	row := db.QueryRow(sqlCount)
	err = row.Scan(&result.Count)

	return result, err
}

func CreateOrders(col []string, values []string) (services.Orders, error) {
	sql := fmt.Sprint(`INSERT INTO "orders" (`, strings.Join(col, ", "), `) VALUES (`, strings.Join(values, ", "), `) RETURNING *`)
	data := services.Orders{}
	err := db.Get(&data, sql)
	return data, err
}

func UpdateOrders(values []string, id int) (services.Orders, error) {
	sql := fmt.Sprint(`UPDATE "orders" SET`, strings.Join(values, ", "), `, "updatedAt"=now() WHERE "id"=$1 RETURNING *`)
	data := services.Orders{}
	err := db.Get(&data, sql, id)
	return data, err
}

func UpdateOrdersCS(id int, total int, tax int) (services.Orders, error) {
	sql := `UPDATE "orders" SET "total" = $2, "taxAmount" = $3 WHERE "id"=$1 RETURNING *`
	data := services.Orders{}
	err := db.Get(&data, sql, id, total, tax)
	return data, err
}

func DeleteOrders(id int) (services.Orders, error) {
	sql := `DELETE FROM "orders" WHERE "id" = $1 RETURNING *`
	data := services.Orders{}
	err := db.Get(&data, sql, id)
	return data, err
}
