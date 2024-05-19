package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/api/products", productController.Index)
	r.GET("/api/product/:id", productController.Index)
	r.POST("/api/product", productController.Index)
	r.PUT("/api/product/:id", productController.Index)
	r.DELETE("/api/product", productController.Index)

}
