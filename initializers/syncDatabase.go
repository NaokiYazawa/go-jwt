package initializers

import "github.com/NaokiYazawa/go-jwt/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
