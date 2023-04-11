package initializers

import "github.com/Alfred-mk/go-jwt/models"

func SyncDatabase()  {
	DB.AutoMigrate(&models.User{})
}