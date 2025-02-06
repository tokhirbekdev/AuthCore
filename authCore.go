package authCore

import (
    "AuthCore/config"
    "AuthCore/routes"
    "AuthCore/migrations"
)

func InitAuth() {
    config.ConnectToDatabase()
    migrations.Migrate()
}

func StartServer() {
    router := routes.SetupRouter()
    router.Run(":8080")
}
