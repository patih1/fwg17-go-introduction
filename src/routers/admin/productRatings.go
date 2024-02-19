package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/patih1/fwg17-go-backend/src/controllers/admin"
)

func ProductRatingsRouter(r *gin.RouterGroup) {
	r.GET("", admin.ListAllProductRatings)
	r.GET("/:id", admin.DetailProductRatings)
	r.POST("", admin.CreateProductRatings)
	r.PATCH("/:id", admin.UpdateProductRatings)
	r.DELETE("/:id", admin.DeleteProductRatings)
}
