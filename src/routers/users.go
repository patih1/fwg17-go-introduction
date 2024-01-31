package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/patih1/fwg17-go-backend/src/controllers"
)

func UserRouter(r *gin.RouterGroup) {
	r.GET("", controllers.ListAllUsers)
	r.POST("", controllers.CreateUser)
}
