package customer

import (
	"github.com/gin-gonic/gin"
	"github.com/patih1/fwg17-go-backend/src/controllers/customer"
)

func OrderRouters(r *gin.RouterGroup) {
	r.GET("/", customer.HistoryOrder)
	r.GET("/:id", customer.DetailHistory)
	r.POST("/", customer.CreateOrders)
}
