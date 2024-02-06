package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matthewhartstonge/argon2"
	"github.com/patih1/fwg17-go-backend/src/models"
	"github.com/patih1/fwg17-go-backend/src/services"
)

func Login(c *gin.Context) {
	form := services.User{}
	err := c.ShouldBind(&form)

	if err != nil {
		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
			Success: false,
			Message: "internal server error",
		})
		return
	}

	found, err := models.FindEmail(form.Email)

	if err != nil {
		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
			Success: false,
			Message: "couldn't find email",
		})
		return
	}

	ok, _ := argon2.VerifyEncoded([]byte(form.Password), []byte(found.Password))

	// token, err := middleware.CreateToken(found.Role)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	if ok {
		c.JSON(http.StatusOK, &services.ResponseOnly{
			Success: true,
			Message: "login successfully",
			// Token:   token,
		})
	} else {
		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
			Success: false,
			Message: "wrong password",
		})
		return
	}
}

func Register(c *gin.Context) {
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
	_, err := models.Create(data)

	if err != nil {
		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
			Success: false,
			Message: "email already used",
		})
		return
	}

	c.JSON(http.StatusOK, &services.ResponseOnly{
		Success: true,
		Message: "create user successfully",
	})
}
