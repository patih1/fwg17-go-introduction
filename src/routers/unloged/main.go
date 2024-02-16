package unloged

import (
	"github.com/gin-gonic/gin"
)

func Combine(r *gin.RouterGroup) {
	ProductRouter(r.Group("/product"))
}
