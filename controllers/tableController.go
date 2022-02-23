package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pi-rate14/stove-backend/database"
	"go.mongodb.org/mongo-driver/mongo"
)

var tableCollection *mongo.Collection = database.OpenCollection(database.Client, "table")


func GetTables() gin.HandlerFunc{
	return func(c *gin.Context){

	}
}

func GetTable() gin.HandlerFunc{
	return func(c *gin.Context){

	}
}

func CreateTable() gin.HandlerFunc{
	return func(c *gin.Context){

	}
}

func UpdateTable() gin.HandlerFunc{
	return func(c *gin.Context){

	}
}