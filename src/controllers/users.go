package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PageInfo struct {
	Page int `json:"page"`
}

type Response struct {
	Success  bool        `json:"success"`
	Message  string      `json:"message"`
	Pageinfo PageInfo    `json:"pageInfo"`
	Results  interface{} `json:"results"`
}

type User struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func ListAllUsers(c *gin.Context) {
	// id := c.Param("id")
	page, _ := strconv.Atoi(c.Query("page"))
	// data := map[string]interface{}{
	// 	"id":       id,
	// 	"fullName": "Khairullah Haidar",
	// 	"email":    "haidar@mail.com",
	// }
	c.JSON(http.StatusOK, &Response{
		Success: true,
		Message: "OK",
		Pageinfo: PageInfo{
			Page: page,
		},
		Results: []User{
			{
				Id:       1,
				Email:    "hai@mail.com",
				Password: "1234",
			},
		},
	})
}
