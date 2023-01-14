package configDB

import "github.com/astrotwig/astrocore_2/models"

func ConfigDb() {
	DB.AutoMigrate(&models.User{})
}
