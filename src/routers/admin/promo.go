package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/patih1/fwg17-go-backend/src/controllers/admin"
)

func PromoRouter(r *gin.RouterGroup) {
	r.GET("", admin.ListAllPromo)
	r.GET("/:id", admin.DetailPromo)
	r.POST("", admin.CreatePromo)
	r.PATCH("/:id", admin.UpdatePromo)
	r.DELETE("/:id", admin.DeletePromo)
}
