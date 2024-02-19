package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/patih1/fwg17-go-backend/src/controllers/admin"
)

func ProductCategoriesRouter(r *gin.RouterGroup) {
	r.GET("", admin.ListAllProductCategories)
	r.GET("/:id", admin.DetailProductCategories)
	r.POST("", admin.CreateProductCategories)
	r.PATCH("/:id", admin.UpdateProductCategories)
	r.DELETE("/:id", admin.DeleteProductCategories)
}
