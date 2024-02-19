package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/patih1/fwg17-go-backend/src/controllers/admin"
)

func OrdersRouter(r *gin.RouterGroup) {
	r.GET("", admin.ListAllOrders)
	r.GET("/:id", admin.DetailOrders)
	r.POST("", admin.CreateOrders)
	r.PATCH("/:id", admin.UpdateOrders)
	r.DELETE("/:id", admin.DeleteOrders)
}
