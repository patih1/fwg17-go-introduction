package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/patih1/fwg17-go-backend/src/routers"
	"github.com/patih1/fwg17-go-backend/src/services"
)

func main() {

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{"GET", "POST", "PATCH", "DELETE"},
	}))
	routers.Combine(r)
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, &services.ResponseOnly{
			Success: false,
			Message: "Resource not found restart on changes",
		})
	})
	r.Run(":8888")
	// r.Run("127.0.0.1:8888")
}
