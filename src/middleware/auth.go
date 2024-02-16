package middlewares

import (
	"errors"
	"net/http"
	"strings"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/matthewhartstonge/argon2"
	"github.com/patih1/fwg17-go-backend/src/models"
	"github.com/patih1/fwg17-go-backend/src/services"
)

func Auth() (*jwt.GinJWTMiddleware, error) {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "go-backend",
		Key:         []byte("secret"),
		IdentityKey: "id",
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			user := data.(services.User)
			return jwt.MapClaims{
				"id":   user.Id,
				"role": user.Role,
			}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			roleStr := claims["role"].(string)
			return &services.Login{
				Id:   int(claims["id"].(float64)),
				Role: &roleStr,
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			form := services.Login{}
			err := c.ShouldBind(&form)

			if err != nil {
				return nil, err
			}

			found, err := models.FindEmail(form.Email)

			if err != nil {
				return nil, err
			}

			ok, _ := argon2.VerifyEncoded([]byte(*form.Password), []byte(*found.Password))

			if ok {
				return found, nil
			} else {
				return nil, errors.New("invalid password")
			}
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			user := data.(*services.Login)
			if strings.HasPrefix(c.Request.URL.Path, "/admin") {
				if *user.Role != "admin" {
					return false
				}
			} else if strings.HasPrefix(c.Request.URL.Path, "/customer") {
				if *user.Role != "customer" && *user.Role != "admin" {
					return false
				}
			}
			return true
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(http.StatusUnauthorized, &services.ResponseOnly{
				Success: false,
				Message: "unauthorized",
			})
		},
		LoginResponse: func(c *gin.Context, code int, token string, time time.Time) {
			c.JSON(http.StatusOK, &services.Response{
				Success: true,
				Message: "login success",
				Results: struct {
					Token string
				}{
					Token: token,
				},
			})

		},
	})
	if err != nil {
		return nil, err
	}
	return authMiddleware, nil
}
