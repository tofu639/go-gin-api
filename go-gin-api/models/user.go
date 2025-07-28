package models

import (
    "time"
)

// User represents a user entity
type User struct {
    ID        uint      `json:"id" example:"1"`
    Name      string    `json:"name" example:"John Doe"`
    Email     string    `json:"email" example:"john@example.com"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

// CreateUserRequest represents the request body for creating a user
type CreateUserRequest struct {
    Name  string `json:"name" binding:"required" example:"John Doe"`
    Email string `json:"email" binding:"required,email" example:"john@example.com"`
}

// UpdateUserRequest represents the request body for updating a user
type UpdateUserRequest struct {
    Name  string `json:"name" example:"John Doe"`
    Email string `json:"email" example:"john@example.com"`
}