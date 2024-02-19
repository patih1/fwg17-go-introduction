package admin

import (
	"github.com/gin-gonic/gin"
	middlewares "github.com/patih1/fwg17-go-backend/src/middleware"
)

func Combine(r *gin.RouterGroup) {
	authMiddleware, _ := middlewares.Auth()
	r.Use(authMiddleware.MiddlewareFunc())
	UserRouter(r.Group("/user"))
	ProductRouter(r.Group("/product"))
	MessageRouter(r.Group("/message"))
	OrderDetailsRouter(r.Group("/order-detail"))
	OrdersRouter(r.Group("/order"))
	ProductCategoriesRouter(r.Group("/product-category"))
	ProductRatingsRouter(r.Group("/product-rating"))
	ProductSizeRouter(r.Group("/product-size"))
	ProductTagsRouter(r.Group("/product-tag"))
	ProductVariantRouter(r.Group("/product-variant"))
	PromoRouter(r.Group("/promo"))
	TagsRouter(r.Group("/tag"))
}
