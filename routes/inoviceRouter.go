package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/pi-rate14/stove-backend/controllers"
)

func InoviceRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/inovices", controller.GetInovices())
	incomingRoutes.GET("/inovices/:inovice_id", controller.GetInovice())
	incomingRoutes.POST("/inovices", controller.CreateInovice())
	incomingRoutes.PATCH("/inovices/:inovice_id", controller.UpdateInovice())
}