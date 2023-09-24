package api

import (
	"github.com/gin-gonic/gin"
	_ "github.com/khasmag06/zdravservice-test/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	*gin.Engine
	productService productsService
	logger         logger
}

func NewHandler(ps productsService, l logger) *Handler {
	h := &Handler{
		Engine:         gin.New(),
		productService: ps,
		logger:         l,
	}

	h.Use(gin.Recovery())

	// Swagger
	h.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := h.Group("/api")

	api.GET("products/get", h.getProducts)

	return h

}
