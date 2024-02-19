package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/patih1/fwg17-go-backend/src/controllers/admin"
)

func ProductVariantRouter(r *gin.RouterGroup) {
	r.GET("", admin.ListAllProductVariant)
	r.GET("/:id", admin.DetailProductVariant)
	r.POST("", admin.CreateProductVariant)
	r.PATCH("/:id", admin.UpdateProductVariant)
	r.DELETE("/:id", admin.DeleteProductVariant)
}
