package models

import (
	"fmt"
	"strings"

	"github.com/patih1/fwg17-go-backend/src/services"
)

func FindAllOrderDetails(limit int, offset int) (services.Info, error) {
	sql := `SELECT * FROM "orderDetails"
	ORDER BY "id" ASC LIMIT $1 OFFSET $2`
	sqlCount := `SELECT COUNT(*) FROM "orderDetails"`
	result := services.Info{}
	data := []services.OrderDetails{}
	err := db.Select(&data, sql, limit, offset)
	result.Data = data

	fmt.Print(err)

	row := db.QueryRow(sqlCount)
	err = row.Scan(&result.Count)

	return result, err
}

func FindOneOrderDetails(id *int) (services.OrderDetails, error) {
	sql := `SELECT * FROM "orderDetails" WHERE
	"id" = $1`
	data := services.OrderDetails{}
	err := db.Get(&data, sql, id)
	return data, err
}

func FindHistoryOD(id *int) (services.OrderDetails, error) {
	sql := `SELECT * FROM "orderDetails" WHERE
	"orderId" = $1`
	data := services.OrderDetails{}
	err := db.Get(&data, sql, id)
	return data, err
}

func CreateOrderDetails(col []string, values []string) (services.OrderDetails, error) {
	sql := fmt.Sprint(`INSERT INTO "orderDetails" (`, strings.Join(col, ", "), `) VALUES (`, strings.Join(values, ", "), `) RETURNING *`)
	data := services.OrderDetails{}
	err := db.Get(&data, sql)
	return data, err
}

func UpdateOrderDetails(values []string, id int) (services.OrderDetails, error) {
	sql := fmt.Sprint(`UPDATE "orderDetails" SET`, strings.Join(values, ", "), `, "updatedAt"=now() WHERE "id"=$1 RETURNING *`)
	data := services.OrderDetails{}
	err := db.Get(&data, sql, id)
	return data, err
}

func DeleteOrderDetails(id int) (services.OrderDetails, error) {
	sql := `DELETE FROM "orderDetails" WHERE "id" = $1 RETURNING *`
	data := services.OrderDetails{}
	err := db.Get(&data, sql, id)
	return data, err
}
