package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/patih1/fwg17-go-backend/src/controllers"
	middlewares "github.com/patih1/fwg17-go-backend/src/middleware"
)

func AuthRouter(r *gin.RouterGroup) {
	authmiddleware, _ := middlewares.Auth()
	r.POST("/login", authmiddleware.LoginHandler)
	r.POST("/register", controllers.Register)
	r.POST("/forgot-password", controllers.ForgotPassword)
}
