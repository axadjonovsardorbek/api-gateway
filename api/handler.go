package api

import (
	// "github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	swaggerFiles "github.com/swaggo/files"

	ginSwagger "github.com/swaggo/gin-swagger"

	_ "api-gateway/api/docs"
	"api-gateway/api/handlers"
	"api-gateway/api/middleware"
	"api-gateway/config/logger"
)

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func NewRouter(connR, connP *grpc.ClientConn, logger logger.Logger) *gin.Engine {
	h := handlers.NewHandler(connR, connP, logger)
	router := gin.Default()

	router.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	protected := router.Group("/", middleware.JWTMiddleware())

	reservation := protected.Group("/reservation")
	reservation.POST("/", h.ReservationCreate)
	reservation.GET("/:id", h.ReservationGet)
	reservation.PUT("/:id", h.ReservationUpdate)
	reservation.DELETE("/:id", h.ReservationDelete)
	protected.GET("/reservations", h.ReservationGetAll)

	return router
}
