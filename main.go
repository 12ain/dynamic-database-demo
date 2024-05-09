package main

import (
	"dynamic-database-demo/config"
	"dynamic-database-demo/handlers"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"strconv"
)

func main() {
	// read config
	config.InitConfig()

	// read gin config
	ginMode := viper.GetString("gin.mode")
	ginPort := viper.GetInt("gin.port")

	r := gin.Default()

	// set gin mode
	if ginMode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// init repository
	handlers.InitRepo()

	// register routes
	r.POST("/users", handlers.CreateUser)
	r.GET("/users/:id", handlers.GetUser)
	r.PUT("/users/:id", handlers.UpdateUser)
	r.DELETE("/users/:id", handlers.DeleteUser)

	r.Run("0.0.0.0:" + strconv.Itoa(ginPort))

}
