package customer

import (
	"github.com/gin-gonic/gin"
	"github.com/patih1/fwg17-go-backend/src/controllers/customer"
)

func UserRouter(r *gin.RouterGroup) {
	r.GET("/", customer.DetailUser)
	r.PATCH("/", customer.UpdateUser)
}
