package database

import (
    "log"
    "os"

    "gorm.io/driver/mysql"
    "gorm.io/gorm"

    "github.com/bethojunior/go-blog/internal/models"
)

func InitDB() *gorm.DB {
    dsn := os.Getenv("DB_DSN")
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Erro ao conectar no banco:", err)
    }

    err = db.AutoMigrate(&models.User{})
    if err != nil {
        log.Fatal("Erro ao migrar tabela User:", err)
    }

    return db
}
