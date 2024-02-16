package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/patih1/fwg17-go-backend/src/controllers/admin"
)

func UserRouter(r *gin.RouterGroup) {
	r.GET("", admin.ListAllUsers)
	r.GET("/:id", admin.DetailUser)
	r.POST("", admin.CreateUser)
	r.PATCH("/:id", admin.UpdateUser)
	r.DELETE("/:id", admin.DeleteUser)
}
