package main

import (
    "log"

    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"

    "github.com/bethojunior/go-blog/internal/database"
    "github.com/bethojunior/go-blog/internal/routes"
    "github.com/bethojunior/go-blog/internal/wire"
)

func main() {
    if err := godotenv.Load(); err != nil {
        log.Println("⚠️  .env não carregado, continuando com variáveis do sistema")
    }

    db := database.InitDB()

    userHandler := wire.InitializeUserHandler(db)

    r := gin.Default()

    routes.SetupRoutes(r, userHandler)

    r.Run(":8080")
}
