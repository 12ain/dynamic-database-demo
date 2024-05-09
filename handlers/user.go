package handlers

import (
	"dynamic-database-demo/database"
	"dynamic-database-demo/model"
	"dynamic-database-demo/repository"

	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var userRepository repository.UserRepository // global variable to hold the repository instance

func InitRepo() {
	DBDriver := viper.GetString("database.driver")
	fmt.Println("Using database driver:", DBDriver)

	switch DBDriver {
	case string(database.MongoDB):
		db, err := database.NewMongoDBInstance()
		if err != nil {
			panic(err)
		}
		userRepository = repository.NewMongoDBUserRepository(db)
	case string(database.MySQL):
		db, err := database.NewMySQLInstance()
		if err != nil {
			panic(err)
		}
		userRepository = repository.NewMySQLUserRepository(db)
	default:
		panic("Unsupported database driver")
	}
}

// CreateUser create a new user
func CreateUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := userRepository.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

// GetUser get a user by id
func GetUser(c *gin.Context) {
	userID := c.Param("id")
	user, err := userRepository.GetUser(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// UpdateUser update a user by id
func UpdateUser(c *gin.Context) {
	userID := c.Param("id")
	id, _ := strconv.Atoi(userID)
	var updatedUser model.User
	updatedUser.ID = uint(id)
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := userRepository.UpdateUser(updatedUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

// DeleteUser delete a user by id
func DeleteUser(c *gin.Context) {
	userID := c.Param("id")
	err := userRepository.DeleteUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
