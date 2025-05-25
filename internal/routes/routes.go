package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/bethojunior/go-blog/internal/handlers"
)

func SetupRoutes(r *gin.Engine, userHandler *handlers.UserHandler) {
    userGroup := r.Group("/users")
    {
        userGroup.POST("/", userHandler.CreateUser)
        userGroup.GET("/", userHandler.ListUser)
        userGroup.GET("/:id", userHandler.GetUser)
    }
}

