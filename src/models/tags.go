package models

import (
	"fmt"
	"strings"

	"github.com/patih1/fwg17-go-backend/src/services"
)

func FindAllTags(limit int, offset int) (services.Info, error) {
	sql := `SELECT * FROM "tags"
	ORDER BY "id" ASC LIMIT $1 OFFSET $2`
	sqlCount := `SELECT COUNT(*) FROM "tags"`
	result := services.Info{}
	data := []services.Tags{}
	err := db.Select(&data, sql, limit, offset)
	result.Data = data

	fmt.Print(err)

	row := db.QueryRow(sqlCount)
	err = row.Scan(&result.Count)

	return result, err
}

func FindOneTags(id *int) (services.Tags, error) {
	sql := `SELECT * FROM "tags" WHERE
	"id" = $1`
	data := services.Tags{}
	err := db.Get(&data, sql, id)
	return data, err
}

func CreateTags(col []string, values []string) (services.Tags, error) {
	sql := fmt.Sprint(`INSERT INTO "tags" (`, strings.Join(col, ", "), `) VALUES (`, strings.Join(values, ", "), `) RETURNING *`)
	data := services.Tags{}
	err := db.Get(&data, sql)
	return data, err
}

func UpdateTags(values []string, id int) (services.Tags, error) {
	sql := fmt.Sprint(`UPDATE "tags" SET`, strings.Join(values, ", "), `, "updatedAt"=now() WHERE "id"=$1 RETURNING *`)
	data := services.Tags{}
	err := db.Get(&data, sql, id)
	return data, err
}

func DeleteTags(id int) (services.Tags, error) {
	sql := `DELETE FROM "tags" WHERE "id" = $1 RETURNING *`
	data := services.Tags{}
	err := db.Get(&data, sql, id)
	return data, err
}
