package models

import (
	"fmt"
	"strings"

	"github.com/patih1/fwg17-go-backend/src/services"
)

func FindAllMessage(limit int, offset int, search string) (services.Info, error) {
	sql := `SELECT * FROM "message"	
	WHERE "text" ILIKE $3 ORDER BY "id" ASC LIMIT $1 OFFSET $2`
	sqlCount := fmt.Sprintf(`SELECT COUNT(*) FROM "message" where "text" ILIKE '%%%v%%'`, search)
	result := services.Info{}
	data := []services.Message{}
	fmtSearch := fmt.Sprintf("%%%v%%", search)
	err := db.Select(&data, sql, limit, offset, fmtSearch)
	result.Data = data

	fmt.Print(err)

	row := db.QueryRow(sqlCount)
	err = row.Scan(&result.Count)

	return result, err
}

func FindUserMessage(limit int, offset int, search string, id int) (services.Info, error) {
	sql := `SELECT "m"."id", "m"."senderId", "m"."recipientId", "m"."text", "u1"."fullName" FROM "message" "m"
	INNER join "users" "u1" on "u1"."id" = "m"."senderId"
	INNER join "users" "u2" on "u2"."id" = "m"."recipientId"
	WHERE "m"."senderId"=$3 or "m"."recipientId"=$3 
	ORDER BY "m"."id" ASC LIMIT $1 OFFSET $2`
	sqlCount := fmt.Sprintf(`SELECT COUNT(*) FROM "message" WHERE "senderId"=%v or "recipientId"=%v`, id, id)
	result := services.Info{}
	data := []services.CSMessage{}
	err := db.Select(&data, sql, limit, offset, id)
	result.Data = data

	fmt.Print(err)

	row := db.QueryRow(sqlCount)
	err = row.Scan(&result.Count)

	return result, err
}

func FindOneMessage(id *int) (services.Message, error) {
	sql := `SELECT * FROM "message" WHERE
	"id" = $1`
	data := services.Message{}
	err := db.Get(&data, sql, id)
	return data, err
}

func CreateMessage(col []string, values []string) (services.Message, error) {
	sql := fmt.Sprint(`INSERT INTO "message" (`, strings.Join(col, ", "), `) VALUES (`, strings.Join(values, ", "), `) RETURNING *`)
	data := services.Message{}
	err := db.Get(&data, sql)
	return data, err
}

func UpdateMessage(values []string, id int) (services.Message, error) {
	sql := fmt.Sprint(`UPDATE "message" SET`, strings.Join(values, ", "), `, "updatedAt"=now() WHERE "id"=$1 RETURNING *`)
	data := services.Message{}
	err := db.Get(&data, sql, id)
	return data, err
}

func DeleteMessage(id int) (services.Message, error) {
	sql := `DELETE FROM "message" WHERE "id" = $1 RETURNING *`
	data := services.Message{}
	err := db.Get(&data, sql, id)
	return data, err
}
