package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/patih1/fwg17-go-backend/src/controllers/admin"
)

func ProductTagsRouter(r *gin.RouterGroup) {
	r.GET("", admin.ListAllProductTags)
	r.GET("/:id", admin.DetailProductTags)
	r.POST("", admin.CreateProductTags)
	r.PATCH("/:id", admin.UpdateProductTags)
	r.DELETE("/:id", admin.DeleteProductTags)
}
