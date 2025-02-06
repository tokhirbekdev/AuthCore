package authCore

import (
    "github.com/tokhirbekdev/AuthCore/config"
    "github.com/tokhirbekdev/AuthCore/routes"
    "github.com/tokhirbekdev/AuthCore/migrations"
)

func InitAuth() {
    config.ConnectToDatabase()
    migrations.Migrate()
}

func StartServer() {
    router := routes.SetupRouter()
    router.Run(":8080")
}
