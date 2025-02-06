package migrations

import (
    "github.com/tokhirbekdev/AuthCore/config"
    "github.com/tokhirbekdev/AuthCore/models"
)

func Migrate() {
    config.DB.AutoMigrate(&models.User{})
}
