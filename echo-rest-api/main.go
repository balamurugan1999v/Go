package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// User represents the resource model with validation tags
type User struct {
	ID    string `json:"id"`
	Name  string `json:"name" validate:"required,min=2"`
	Email string `json:"email" validate:"required,email"`
}

// CustomValidator integrates go-playground/validator with Echo
type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

// In-memory database
var users = map[string]User{
	"1": {ID: "1", Name: "Alice", Email: "alice@example.com"},
	"2": {ID: "2", Name: "Bob", Email: "bob@example.com"},
}

// Handlers
func getUsers(c echo.Context) error {
	userList := make([]User, 0, len(users))
	for _, u := range users {
		userList = append(userList, u)
	}
	return c.JSON(http.StatusOK, userList)
}

func getUserByID(c echo.Context) error {
	id := c.Param("id")
	user, exists := users[id]
	if !exists {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "User not found"})
	}
	return c.JSON(http.StatusOK, user)
}

func createUser(c echo.Context) error {
	var user User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request payload"})
	}

	if err := c.Validate(&user); err != nil {
		return err
	}

	if _, exists := users[user.ID]; exists {
		return c.JSON(http.StatusConflict, echo.Map{"message": "User ID already exists"})
	}

	users[user.ID] = user
	return c.JSON(http.StatusCreated, user)
}

func updateUser(c echo.Context) error {
	id := c.Param("id")
	if _, exists := users[id]; !exists {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "User not found"})
	}

	var updatedUser User
	if err := c.Bind(&updatedUser); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid payload"})
	}

	updatedUser.ID = id
	if err := c.Validate(&updatedUser); err != nil {
		return err
	}

	users[id] = updatedUser
	return c.JSON(http.StatusOK, updatedUser)
}

func deleteUser(c echo.Context) error {
	id := c.Param("id")
	if _, exists := users[id]; !exists {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "User not found"})
	}

	delete(users, id)
	return c.NoContent(http.StatusNoContent)
}

func main() {
	e := echo.New()

	// Register Custom Validator
	e.Validator = &CustomValidator{validator: validator.New()}

	// Middleware
	e.Use(middleware.Logger())  // Log HTTP requests
	e.Use(middleware.Recover()) // Recover from panics safely

	// Routes
	api := e.Group("/api/v1")
	api.GET("/users", getUsers)
	api.GET("/users/:id", getUserByID)
	api.POST("/users", createUser)
	api.PUT("/users/:id", updateUser)
	api.DELETE("/users/:id", deleteUser)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
