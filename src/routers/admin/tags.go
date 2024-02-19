package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/patih1/fwg17-go-backend/src/controllers/admin"
)

func TagsRouter(r *gin.RouterGroup) {
	r.GET("", admin.ListAllTags)
	r.GET("/:id", admin.DetailTags)
	r.POST("", admin.CreateTags)
	r.PATCH("/:id", admin.UpdateTags)
	r.DELETE("/:id", admin.DeleteTags)
}
