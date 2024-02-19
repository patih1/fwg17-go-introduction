package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/patih1/fwg17-go-backend/src/controllers/admin"
)

func MessageRouter(r *gin.RouterGroup) {
	r.GET("", admin.ListAllMessage)
	r.GET("/:id", admin.DetailMessage)
	r.POST("", admin.CreateMessage)
	r.PATCH("/:id", admin.UpdateMessage)
	r.DELETE("/:id", admin.DeleteMessage)
}
