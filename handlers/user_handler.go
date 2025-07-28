package handlers

import (
    "net/http"
    "strconv"
    "time"

    "go-gin-api/models"

    "github.com/gin-gonic/gin"
)

// In-memory storage (replace with database in production)
var users = []models.User{
    {ID: 1, Name: "Alice Johnson", Email: "alice@example.com", CreatedAt: time.Now(), UpdatedAt: time.Now()},
    {ID: 2, Name: "Bob Smith", Email: "bob@example.com", CreatedAt: time.Now(), UpdatedAt: time.Now()},
}

var nextID uint = 3

// GetUser godoc
// @Summary Get a user by ID
// @Description Get a user from the database by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} models.User
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /users/{id} [get]
func GetUser(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    for _, user := range users {
        if user.ID == uint(id) {
            c.JSON(http.StatusOK, user)
            return
        }
    }

    c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}

// GetUsers godoc
// @Summary Get all users
// @Description Get all users from the database
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {array} models.User
// @Router /users [get]
func GetUsers(c *gin.Context) {
    c.JSON(http.StatusOK, users)
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with the provided information
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.CreateUserRequest true "User information"
// @Success 201 {object} models.User
// @Failure 400 {object} map[string]string
// @Router /users [post]
func CreateUser(c *gin.Context) {
    var req models.CreateUserRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Check if email already exists
    for _, user := range users {
        if user.Email == req.Email {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
            return
        }
    }

    newUser := models.User{
        ID:        nextID,
        Name:      req.Name,
        Email:     req.Email,
        CreatedAt: time.Now(),
        UpdatedAt: time.Now(),
    }

    users = append(users, newUser)
    nextID++

    c.JSON(http.StatusCreated, newUser)
}

// UpdateUser godoc
// @Summary Update a user
// @Description Update an existing user with the provided information
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body models.UpdateUserRequest true "User information"
// @Success 200 {object} models.User
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /users/{id} [put]
func UpdateUser(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    var req models.UpdateUserRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    for i, user := range users {
        if user.ID == uint(id) {
            if req.Name != "" {
                users[i].Name = req.Name
            }
            if req.Email != "" {
                // Check if email already exists (excluding current user)
                for _, existingUser := range users {
                    if existingUser.ID != uint(id) && existingUser.Email == req.Email {
                        c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
                        return
                    }
                }
                users[i].Email = req.Email
            }
            users[i].UpdatedAt = time.Now()
            c.JSON(http.StatusOK, users[i])
            return
        }
    }

    c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}

// DeleteUser godoc
// @Summary Delete a user
// @Description Delete a user from the database by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /users/{id} [delete]
func DeleteUser(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    for i, user := range users {
        if user.ID == uint(id) {
            users = append(users[:i], users[i+1:]...)
            c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
            return
        }
    }

    c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}