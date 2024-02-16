package models

import (
	"fmt"
	"time"
)

type FormReset struct {
	Id        int        `db:"id"`
	Email     string     `db:"email" form:"email"`
	Otp       string     `db:"otp" form:"otp"`
	CreatedAt time.Time  `db:"createdAt"`
	UpdatedAt *time.Time `db:"updatedAt"`
}

func FindOtp(otp *string) (FormReset, error) {
	sql := `SELECT * FROM "resetPassword" WHERE "otp" = $1`
	data := FormReset{}
	err := db.Get(&data, sql, otp)
	return data, err
}

func CreateResetPassword(data FormReset) (FormReset, error) {
	fmt.Println(data)
	sql := `INSERT INTO "resetPassword" ("email", "otp") VALUES ( :email, :otp
		) RETURNING *`
	result := FormReset{}
	rows, err := db.NamedQuery(sql, data)

	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}
	return result, err
}

func DeleteResetPassword(id int) (FormReset, error) {
	sql := `DELETE FROM "resetPassword" WHERE "id" = $1 RETURNING *`
	data := FormReset{}
	err := db.Get(&data, sql, id)
	return data, err
}
