package main

import (
	"os"

	"github.com/gin-gonic/gin"
	database "github.com/pi-rate14/stove-backend/database"
	middleware "github.com/pi-rate14/stove-backend/middleware"
	routes "github.com/pi-rate14/stove-backend/routes"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.TableRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemRoutes(router)
	routes.InoviceRoutes(router)

	router.Run(":"+port)
}