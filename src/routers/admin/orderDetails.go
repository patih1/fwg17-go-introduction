package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/patih1/fwg17-go-backend/src/controllers/admin"
)

func OrderDetailsRouter(r *gin.RouterGroup) {
	r.GET("", admin.ListAllOrderDetails)
	r.GET("/:id", admin.DetailOrderDetails)
	r.POST("", admin.CreateOrderDetails)
	r.PATCH("/:id", admin.UpdateOrderDetails)
	r.DELETE("/:id", admin.DeleteOrderDetails)
}
