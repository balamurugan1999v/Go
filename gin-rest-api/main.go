package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Item represents a generic data structure for our REST API
type Item struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// In-memory mock database
var items = []Item{
	{ID: "1", Name: "Laptop", Price: 999.99},
	{ID: "2", Name: "Mouse", Price: 24.50},
	{ID: "3", Name: "Keyboard", Price: 75.00},
}

// GET /items - Retrieve all items
func getItems(c *gin.Context) {
	c.JSON(http.StatusOK, items)
}

// GET /items/:id - Retrieve a specific item by ID
func getItemByID(c *gin.Context) {
	id := c.Param("id")

	for _, item := range items {
		if item.ID == id {
			c.JSON(http.StatusOK, item)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "item not found"})
}

// POST /items - Add a new item
func createItem(c *gin.Context) {
	var newItem Item

	// Bind incoming JSON body to the newItem struct
	if err := c.ShouldBindJSON(&newItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	items = append(items, newItem)
	c.JSON(http.StatusCreated, newItem)
}

// PUT /items/:id - Update an existing item
func updateItem(c *gin.Context) {
	id := c.Param("id")

	var updatedItem Item
	if err := c.ShouldBindJSON(&updatedItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, item := range items {
		if item.ID == id {
			items[i].Name = updatedItem.Name
			items[i].Price = updatedItem.Price
			c.JSON(http.StatusOK, items[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "item not found"})
}

// DELETE /items/:id - Delete an item by ID
func deleteItem(c *gin.Context) {
	id := c.Param("id")

	for i, item := range items {
		if item.ID == id {
			// Remove the item from the slice
			items = append(items[:i], items[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "item deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "item not found"})
}

func main() {
	// Initialize default Gin router with Logger and Recovery middleware
	router := gin.Default()

	// Define routes
	router.GET("/items", getItems)
	router.GET("/items/:id", getItemByID)
	router.POST("/items", createItem)
	router.PUT("/items/:id", updateItem)
	router.DELETE("/items/:id", deleteItem)

	// Run the server on localhost:8080
	router.Run("localhost:8080")
}
