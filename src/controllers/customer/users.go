package customer

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/matthewhartstonge/argon2"
	"github.com/patih1/fwg17-go-backend/src/lib"
	middlewares "github.com/patih1/fwg17-go-backend/src/middleware"
	"github.com/patih1/fwg17-go-backend/src/models"
	"github.com/patih1/fwg17-go-backend/src/services"
)

func DetailUser(c *gin.Context) {

	// Get jwt payload using Identity handler
	x, _ := middlewares.Auth()
	data := x.IdentityHandler(c).(*services.Login)

	user, err := models.FindOne(&data.Id)

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

func UpdateUser(c *gin.Context) {

	// Get jwt payload using Identity handler
	x, _ := middlewares.Auth()
	loged := x.IdentityHandler(c).(*services.Login)

	data := services.User{}

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
		if x != "Id" && x != "Role" && x != "CreatedAt" && x != "UpdatedAt" {

			// check for empty field
			if !val.Elem().Field(i).IsNil() {
				y := fmt.Sprint(`"`, strings.ToLower(x[:1])+x[1:], `"='`, val.Elem().FieldByName(x).Elem().Interface(), `'`)
				values = append(values, y)
			}
		}
	}
	user, err := models.DynamicUpdate(values, loged.Id)

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
