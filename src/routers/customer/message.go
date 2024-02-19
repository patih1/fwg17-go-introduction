package customer

import (
	"github.com/gin-gonic/gin"
	"github.com/patih1/fwg17-go-backend/src/controllers/customer"
)

func MessageRouters(r *gin.RouterGroup) {
	r.GET("/", customer.ListAllMessage)
	// r.GET("/:id", customer.DetailHistory)
	r.POST("/", customer.CreateMessage)
}
