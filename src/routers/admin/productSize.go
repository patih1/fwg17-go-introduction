package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/patih1/fwg17-go-backend/src/controllers/admin"
)

func ProductSizeRouter(r *gin.RouterGroup) {
	r.GET("", admin.ListAllProductSize)
	r.GET("/:id", admin.DetailProductSize)
	r.POST("", admin.CreateProductSize)
	r.PATCH("/:id", admin.UpdateProductSize)
	r.DELETE("/:id", admin.DeleteProductSize)
}
