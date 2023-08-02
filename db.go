package main

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "log"
    "os"

    "github.com/joho/godotenv"
)

func ConnectDB() *gorm.DB {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    dbUsername := os.Getenv("DB_USERNAME")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbName := os.Getenv("DB_NAME")

    dsn := dbUsername + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"

    var errDB error
    db, errDB = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if errDB != nil {
        log.Fatal("Failed to connect to database")
    }

    log.Println("Connected to database")
    return db
}