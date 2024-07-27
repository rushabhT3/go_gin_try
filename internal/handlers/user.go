// internal/handlers/user.go
package handlers

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "github.com/rushabhT3/go_gin_try/internal/models"
)

func GetUsers(c *gin.Context) {
    c.JSON(http.StatusOK, models.Users)
}

func GetUser(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    for _, user := range models.Users {
        if user.ID == id {
            c.JSON(http.StatusOK, user)
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
}

func CreateUser(c *gin.Context) {
    var newUser models.User
    if err := c.ShouldBindJSON(&newUser); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    newUser.ID = len(models.Users) + 1
    models.Users = append(models.Users, newUser)
    c.JSON(http.StatusCreated, newUser)
}

func UpdateUser(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var updatedUser models.User
    if err := c.ShouldBindJSON(&updatedUser); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    for i, user := range models.Users {
        if user.ID == id {
            models.Users[i] = updatedUser
            models.Users[i].ID = id
            c.JSON(http.StatusOK, models.Users[i])
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
}

func DeleteUser(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    for i, user := range models.Users {
        if user.ID == id {
            models.Users = append(models.Users[:i], models.Users[i+1:]...)
            c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
}
