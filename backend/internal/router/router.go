package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/user/controle-estoque/backend/internal/handler"
)

func SetupRouter(productHandler *handler.ProductHandler, movementHandler *handler.MovementHandler) *gin.Engine {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	api := r.Group("/api/v1")
	{
		products := api.Group("/products")
		{
			products.POST("", productHandler.Create)
			products.GET("", productHandler.List)
			products.GET("/:id", productHandler.GetByID)
			products.PUT("/:id", productHandler.Update)
			products.DELETE("/:id", productHandler.Delete)
		}

		movements := api.Group("/movements")
		{
			movements.POST("", movementHandler.Register)
			movements.GET("", movementHandler.List)
		}
	}

	return r
}
