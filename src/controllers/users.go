package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/patih1/fwg17-go-backend/src/models"
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

type ResponseOnly struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
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
			}, {
				Id:       2,
				Email:    "haid@mail.com",
				Password: "123",
			},
		},
	})
}

func CreateUser(c *gin.Context) {
	data := models.User{}

	c.Bind(data)

	user, err := models.CreateUser(data)

	if err != nil {
		log.Fatalln()
		c.JSON(http.StatusInternalServerError, &ResponseOnly{
			Success: false,
			Message: "internal server error",
		})
		fmt.Println("asw")
	}

	c.JSON(http.StatusOK, &Response{
		Success: true,
		Message: "create user successfully",
		Results: user,
	})
}
