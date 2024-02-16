package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/patih1/fwg17-go-backend/src/controllers/admin"
)

func ProductRouter(r *gin.RouterGroup) {
	r.GET("", admin.ListAllProduct)
	r.GET("/:id", admin.DetailProduct)
	r.POST("", admin.CreateProduct)
	r.PATCH("/:id", admin.UpdateProduct)
	r.DELETE("/:id", admin.DeleteProduct)
}
