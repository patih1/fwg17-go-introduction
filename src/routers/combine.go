package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/patih1/fwg17-go-backend/src/routers/admin"
	"github.com/patih1/fwg17-go-backend/src/routers/customer"
	"github.com/patih1/fwg17-go-backend/src/routers/unloged"
)

func Combine(r *gin.Engine) {
	unloged.Combine(r.Group(""))
	AuthRouter(r.Group("/auth"))
	admin.Combine(r.Group("/admin"))
	customer.Combine(r.Group("/customer"))
}
