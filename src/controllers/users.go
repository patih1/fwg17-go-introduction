package controllers

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/matthewhartstonge/argon2"
	"github.com/patih1/fwg17-go-backend/src/models"
	"github.com/patih1/fwg17-go-backend/src/services"
)

func ListAllUsers(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "5"))
	search := c.DefaultQuery("search", "")
	// sortBy := c.DefaultQuery("sortBy", "id")
	offset := (page - 1) * limit
	result, err := models.FindAll(limit, offset, search)

	pageInfo := services.PageInfo{
		Page:      page,
		Limit:     limit,
		LastPage:  int(math.Ceil(float64(result.Count) / float64(limit))),
		TotalData: result.Count,
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
			Success: false,
			Message: "internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, &services.Response{
		Success:  true,
		Message:  "list all users",
		Pageinfo: pageInfo,
		Results:  result.Data,
	})
}

func DetailUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := models.FindOne(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
			Success: false,
			Message: "internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, &services.Response{
		Success: true,
		Message: "detail user",
		Results: user,
	})
}

func CreateUser(c *gin.Context) {
	data := services.User{}

	argon := argon2.DefaultConfig()

	c.ShouldBind(&data)

	// fmt.Println(data.Password)

	if data.Password == "" {
		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
			Success: false,
			Message: "Password cannot be empty",
		})
		return
	}

	encoded, _ := argon.HashEncoded([]byte(data.Password))
	data.Password = string(encoded)
	user, err := models.Create(data)

	if err != nil {
		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
			Success: false,
			Message: "email already existed",
		})
		return
	}

	c.JSON(http.StatusOK, &services.Response{
		Success: true,
		Message: "create user successfully",
		Results: user,
	})
}

func UpdateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data := services.ToUpdateUser{}

	argon := argon2.DefaultConfig()

	c.ShouldBind(&data)
	fmt.Println(reflect.TypeOf(data))
	data.Id = id

	if data.Password != "" {
		encoded, _ := argon.HashEncoded([]byte(data.Password))
		data.Password = string(encoded)
	}

	user, err := models.Update(data)

	if err != nil {
		// log.Fatalln(err)
		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
			Success: false,
			Message: "internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, &services.Response{
		Success: true,
		Message: "update user successfully",
		Results: user,
	})

}

func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := models.Delete(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
			Success: false,
			Message: "internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, &services.Response{
		Success: true,
		Message: "delete user successfully",
		Results: user,
	})
}

func DynamicCreateUser(c *gin.Context) {
	data := services.User{}

	argon := argon2.DefaultConfig()

	c.ShouldBind(&data)

	var col []string
	var values []string

	// fmt.Println(data.Password)

	if data.Password == "" {
		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
			Success: false,
			Message: "Password cannot be empty",
		})
		return
	}

	encoded, _ := argon.HashEncoded([]byte(data.Password))
	data.Password = string(encoded)

	val := reflect.ValueOf(data)
	types := val.Type()

	for i := 0; i < val.NumField(); i++ {
		if types.Field(i).Name != "Id" {

			col = append(col, fmt.Sprint(`"`, strings.ToLower(types.Field(i).Name[:1])+types.Field(i).Name[1:]), `"`)
			values = append(values, fmt.Sprint(`'`, types.Field(i)), `'`)
			fmt.Println(types.Field(i))
		}

	}
	user, err := models.DynamicCreate(col, values)

	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
			Success: false,
			Message: "internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, &services.Response{
		Success: true,
		Message: "create user successfully",
		Results: user,
	})
}
