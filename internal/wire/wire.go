//go:build wireinject
// +build wireinject

package wire

import (
    "github.com/bethojunior/go-blog/internal/handlers"
    "github.com/bethojunior/go-blog/internal/repository"
    "github.com/bethojunior/go-blog/internal/services"
    "gorm.io/gorm"

    "github.com/google/wire"
)

func InitializeUserHandler(db *gorm.DB) *handlers.UserHandler {
    wire.Build(
        repository.NewUserRepository,
        services.NewUserService,
        handlers.NewUserHandler,
    )
    return &handlers.UserHandler{}
}
