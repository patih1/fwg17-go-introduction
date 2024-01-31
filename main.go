package main

import (
	"github.com/gin-gonic/gin"
	"github.com/patih1/fwg17-go-backend/src/routers"
)

func main() {
	r := gin.Default()
	routers.Combine(r)
	r.Run()
}
