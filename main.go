package main

import (
	"github.com/astrotwig/astrocore_2/configDB"
	"github.com/astrotwig/astrocore_2/initializers"
	"github.com/gin-gonic/gin"
)

func init() {

	initializers.LoadEnvVariables()
	configDB.ConnectToDb()
	configDB.ConfigDb()
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Working Perfectly",
		})
	})

	r.Run()

}
