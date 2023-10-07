package configs

import (
    "docker/model"
    "fmt"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "log"
)

var (
    DB *gorm.DB
)

func init() {
    InitDB()
    InitialMigration()
}

func InitDB() {
    connectionString := "root:password@tcp(34.123.248.159:3306)/test5?parseTime=true"
    var err error
    DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %s", err.Error())
    }
}

func InitialMigration() {
    DB.AutoMigrate(&model.User{},&model.Book{})
}
