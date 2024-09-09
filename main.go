package main

import (
    "lms-backend/database"
    "lms-backend/routes"
)

func main() {
    database.InitDB()
    router := routes.SetupRouter()
    router.Run(":8080")
}
