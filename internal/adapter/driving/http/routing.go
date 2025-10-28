package http

import (
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"

	_ "github.com/NekruzRakhimov/product_service/api/docs"

	"github.com/gin-gonic/gin"
)

func (s *Server) endpoints() {
	s.router.GET("/ping", s.Ping)

	s.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiG := s.router.Group("/api", s.checkUserAuthentication)
	{
		apiG.GET("/products", s.GetAllProducts)
		apiG.GET("/products/:id", s.GetProductByID)
		apiG.POST("/products", s.checkIsAdmin, s.CreateProduct)
		apiG.PUT("/products/:id", s.checkIsAdmin, s.UpdateProductByID)
		apiG.DELETE("/products/:id", s.checkIsAdmin, s.DeleteProductByID)
	}

}

// Ping
// @Summary Health-check
// @Description Проверка сервиса
// @Tags Ping
// @Produce json
// @Success 200 {object} CommonResponse
// @Router /ping [get]
func (s *Server) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"ping": "pong",
	})
}
