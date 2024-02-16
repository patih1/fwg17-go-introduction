package unloged

import (
	"github.com/gin-gonic/gin"
	"github.com/patih1/fwg17-go-backend/src/controllers/unloged"
)

func ProductRouter(r *gin.RouterGroup) {
	r.GET("", unloged.ListAllProduct)
	r.GET("/:id", unloged.DetailProduct)
}
