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
    username := "root"
    host := "34.123.248.159"
    port := "3306"
    name := "test5"

    connectionString := fmt.Sprintf("%s@tcp(%s:%s)/%s?parseTime=true",
        username,
        host,
        port,
        name,
    )
    var err error
    DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %s", err.Error())
    }
}

func InitialMigration() {
    DB.AutoMigrate(&model.User{},&model.Book{})
}
