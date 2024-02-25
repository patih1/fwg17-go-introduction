package admin

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/matthewhartstonge/argon2"
	"github.com/patih1/fwg17-go-backend/src/lib"
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
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
			Success: false,
			Message: "internal server error",
		})
		return
	}

	if pageInfo.Page <= 0 || pageInfo.Page > pageInfo.LastPage {
		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
			Success: false,
			Message: "not found",
		})
		return
	}

	c.JSON(http.StatusOK, &services.ResponseWPage{
		Success:  true,
		Message:  "list all users",
		Pageinfo: pageInfo,
		Results:  result.Data,
	})
}

func DetailUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := models.FindOne(&id)

	if err != nil {
		c.JSON(http.StatusNotFound, &services.ResponseOnly{
			Success: false,
			Message: "user not found",
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

	// Argon Config
	argon := argon2.DefaultConfig()

	c.ShouldBind(&data)

	var col []string
	var values []string

	// check for empty not null column
	if *data.Password == "" {
		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
			Success: false,
			Message: "Password cannot be empty",
		})
		return
	} else if *data.Email == "" {
		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
			Success: false,
			Message: "Email cannot be empty",
		})
		return
	} else if *data.FullName == "" {
		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
			Success: false,
			Message: "Email cannot be empty",
		})
		return
	}

	file, err := lib.Upload(c, "profile")

	if file == "wrong ext" {
		c.JSON(http.StatusUnprocessableEntity, &services.ResponseOnly{
			Success: false,
			Message: err.Error(),
		})
		return
	} else if file == "no file" {
		data.Picture = nil
	} else {
		data.Picture = &file
	}

	// hash password
	encoded, _ := argon.HashEncoded([]byte(*data.Password))
	*data.Password = string(encoded)

	val := reflect.ValueOf(&data)
	types := val.Type()

	// iterates trough struct by index number
	for i := 0; i < val.Elem().NumField(); i++ {
		// format struct key's to a simple string
		x := fmt.Sprintf(`%v`, types.Elem().Field(i).Name)

		// prevent process on unecessary column
		if x != "Id" && x != "CreatedAt" && x != "UpdatedAt" {

			// check for empty field
			if !val.Elem().Field(i).IsNil() {

				// insert "key's" to col and "value" to values
				col = append(col, fmt.Sprint(`"`, strings.ToLower(x[:1])+x[1:], `"`))
				values = append(values, fmt.Sprint(`'`, val.Elem().FieldByName(x).Elem().Interface(), `'`))
			}
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

func UpdateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data := services.User{}

	oldFile, _ := models.FindOne(&id)
	// Argon Config
	argon := argon2.DefaultConfig()

	c.ShouldBind(&data)

	var values []string

	file, err := lib.Upload(c, "profile")

	if file == "wrong ext" {
		c.JSON(http.StatusUnprocessableEntity, &services.ResponseOnly{
			Success: false,
			Message: err.Error(),
		})
		return
	} else if file == "no file" {
		data.Picture = nil
	} else {
		data.Picture = &file
		os.Remove(*oldFile.Picture)
	}

	if data.Password != nil {
		encoded, _ := argon.HashEncoded([]byte(*data.Password))
		*data.Password = string(encoded)
	}

	val := reflect.ValueOf(&data)
	types := val.Type()

	// iterates trough struct by index number
	for i := 0; i < val.Elem().NumField(); i++ {
		// format struct key's to a simple string
		x := fmt.Sprintf(`%v`, types.Elem().Field(i).Name)

		// prevent process on unecessary column
		if x != "Id" && x != "CreatedAt" && x != "UpdatedAt" {

			// check for empty field
			if !val.Elem().Field(i).IsNil() {
				y := fmt.Sprint(`"`, strings.ToLower(x[:1])+x[1:], `"='`, val.Elem().FieldByName(x).Elem().Interface(), `'`)
				values = append(values, y)
			}
		}
	}
	user, err := models.DynamicUpdate(values, id)

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
		Message: "update user successfully",
		Results: user,
	})
}

func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := models.Delete(id)

	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
			Success: false,
			Message: "user not found",
		})
		return
	}

	c.JSON(http.StatusOK, &services.Response{
		Success: true,
		Message: "delete user successfully",
		Results: user,
	})
}
