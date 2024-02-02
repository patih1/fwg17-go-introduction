package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/matthewhartstonge/argon2"
	_ "github.com/matthewhartstonge/argon2"
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

type ResponseOnly struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func ListAllUsers(c *gin.Context) {

	page, _ := strconv.Atoi(c.Query("page"))

	users, err := models.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, &ResponseOnly{
			Success: false,
			Message: "internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, &Response{
		Success: true,
		Message: "list all users",
		Pageinfo: PageInfo{
			Page: page,
		},
		Results: users,
	})
}

func DetailUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := models.FindOne(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, &ResponseOnly{
			Success: false,
			Message: "internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, &Response{
		Success: true,
		Message: "detail user",
		Results: user,
	})
}

func CreateUser(c *gin.Context) {
	data := models.User{}

	argon := argon2.DefaultConfig()

	c.Bind(&data)

	encoded, err := argon.HashEncoded([]byte(data.Password))
	data.Password = string(encoded)

	user, err := models.Create(data)

	if err != nil {
		// log.Fatalln(err)
		c.JSON(http.StatusInternalServerError, &ResponseOnly{
			Success: false,
			Message: "internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, &Response{
		Success: true,
		Message: "create user successfully",
		Results: user,
	})
}

func UpdateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data := models.User{}

	argon := argon2.DefaultConfig()

	c.Bind(&data)
	data.Id = id
	encoded, err := argon.HashEncoded([]byte(data.Password))
	data.Password = string(encoded)

	user, err := models.Update(data)

	if err != nil {
		// log.Fatalln(err)
		c.JSON(http.StatusInternalServerError, &ResponseOnly{
			Success: false,
			Message: "internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, &Response{
		Success: true,
		Message: "update user successfully",
		Results: user,
	})

}

func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := models.Delete(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, &ResponseOnly{
			Success: false,
			Message: "internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, &Response{
		Success: true,
		Message: "delete user successfully",
		Results: user,
	})
}
