package database

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "log"
    "os"
)

func InitDB() *gorm.DB {
    dsn := os.Getenv("DB_DSN")
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Erro ao conectar no banco:", err)
    }
    return db
}
