package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matthewhartstonge/argon2"
	"github.com/patih1/fwg17-go-backend/src/lib"
	middlewares "github.com/patih1/fwg17-go-backend/src/middleware"
	"github.com/patih1/fwg17-go-backend/src/models"
	"github.com/patih1/fwg17-go-backend/src/services"
)

type FormReset struct {
	Email           string `form:"email"`
	Otp             string `form:"otp"`
	Password        string `form:"password"`
	ConfirmPassword string `form:"confirmPassword"`
}

func Login(c *gin.Context) {

}

func Register(c *gin.Context) {
	data := services.User{}

	argon := argon2.DefaultConfig()

	c.ShouldBind(&data)

	// fmt.Println(data.Password)

	if *data.Password == "" {
		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
			Success: false,
			Message: "Password cannot be empty",
		})
		return
	}

	encoded, _ := argon.HashEncoded([]byte(*data.Password))
	*data.Password = string(encoded)
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

func ForgotPassword(c *gin.Context) {
	form := FormReset{}
	c.ShouldBind(&form)
	if form.Email != "" {
		found, _ := models.FindEmail(&form.Email)
		if found.Id != 0 {
			FormReset := models.FormReset{
				Otp:   lib.RandomNumber(6),
				Email: *found.Email,
			}

			fmt.Println(FormReset.Otp)

			models.CreateResetPassword(FormReset)
			c.JSON(http.StatusOK, &services.ResponseOnly{
				Success: true,
				Message: "otp has sent to your email",
			})
			return
		} else {
			c.JSON(http.StatusBadRequest, &services.ResponseOnly{
				Success: false,
				Message: "faild to reset",
			})
			return
		}
	}
	if form.Otp != "" {
		found, _ := models.FindOtp(&form.Otp)
		if found.Id != 0 {
			if form.Password == form.ConfirmPassword {
				foundUser, _ := models.FindEmail(&found.Email)
				data := services.User{
					Id: foundUser.Id,
				}

				if data.Password != nil {
					c.JSON(http.StatusBadRequest, &services.ResponseOnly{
						Success: false,
						Message: "password can't be empty",
					})
					return
				}

				argon := argon2.DefaultConfig()
				encoded, _ := argon.HashEncoded([]byte(form.Password))
				x := string(encoded)
				data.Password = &x

				updated, _ := models.UpdatePassword(data)

				message := fmt.Sprintf("password of %s has been changed", *updated.Email)

				c.JSON(http.StatusOK, &services.ResponseOnly{
					Success: true,
					Message: message,
				})
				models.DeleteResetPassword(found.Id)
				return
			} else {
				c.JSON(http.StatusBadRequest, &services.ResponseOnly{
					Success: false,
					Message: "password doesnt match",
				})
				return
			}
		}
	}
	c.JSON(http.StatusBadRequest, &services.ResponseOnly{
		Success: false,
		Message: "internal server error",
	})
}

func GetLogedId(c *gin.Context) int {
	x, _ := middlewares.Auth()
	data := x.IdentityHandler(c).(*services.Login)
	return data.Id
}
