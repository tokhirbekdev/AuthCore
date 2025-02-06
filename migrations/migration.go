package migrations

import (
    "AuthCore/config"
    "AuthCore/models"
)

func Migrate() {
    config.DB.AutoMigrate(&models.User{})
}
