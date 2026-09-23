package controllers

import "github.com/gin-gonic/gin"

func Routers(server *gin.Engine) {
	server.GET("/health", healthCheck)
	server.GET("/getApi", getAPI)
	server.POST("/save", save)
	server.GET("/getAllVehicles", getAllVehicles)
	server.PUT("/updateVehicle", updateVehicle)
}
