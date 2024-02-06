package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/patih1/fwg17-go-backend/src/controllers"
)

func AuthRouter(r *gin.RouterGroup) {
	r.POST("/login", controllers.Login)
	r.POST("/register", controllers.Register)
}
