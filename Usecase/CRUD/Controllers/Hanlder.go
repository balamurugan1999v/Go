package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Vehicle struct {
	ID    int64
	Name  string
	Model string
	Year  string
}

var vehicles []Vehicle = []Vehicle{}

func save(context *gin.Context) {
	var v Vehicle
	fmt.Println("Before")
	fmt.Println(vehicles)
	context.ShouldBindJSON(&v)
	fmt.Println(v)
	vehicles = append(vehicles, v)
	fmt.Println("After")
	fmt.Println(vehicles)
	context.JSON(http.StatusOK, gin.H{"messaage": "Successfully added"})
}

func getAllVehicles(context *gin.Context) {
	context.JSON(http.StatusOK, vehicles)
}

func updateVehicle(context *gin.Context) {
	var v Vehicle
	context.ShouldBindJSON(&v)
	var targetId = v.ID
	for i := range vehicles {
		if vehicles[i].ID == targetId {
			vehicles[i].Model = v.Model
			vehicles[i].Name = v.Name
			vehicles[i].Year = v.Year
		}
	}
	context.JSON(http.StatusOK, gin.H{"message": "Updated successfully"})
}

func healthCheck(context *gin.Context) {
	context.String(http.StatusOK, "Application up and Running")
}

func getAPI(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Application able to server now"})
}

func deleteVehicle(context *gin.Context) {
	var v Vehicle
	context.ShouldBindJSON(&v)
	for i, value := range vehicles {
		if v.ID == value.ID {
			vehicles = append(vehicles[:i], vehicles[i+1:]...)
		}
	}
	context.JSON(http.StatusOK, gin.H{"message": "Vehicle deleted successfully!"})
}
