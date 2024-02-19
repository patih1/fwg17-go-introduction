package customer

import (
	"github.com/gin-gonic/gin"
	middlewares "github.com/patih1/fwg17-go-backend/src/middleware"
)

func Combine(r *gin.RouterGroup) {
	authMiddleware, _ := middlewares.Auth()
	r.Use(authMiddleware.MiddlewareFunc())
	UserRouter(r.Group("/user"))
	OrderRouters(r.Group("/order"))
	MessageRouters(r.Group("/message"))
}
