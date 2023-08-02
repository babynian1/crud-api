

package main

import (
    "github.com/gin-gonic/gin"
    "crud-api/controllers"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "log"
    "os"

    "github.com/joho/godotenv"
)

var db *gorm.DB

func main() {
    router := gin.Default()

    // Menghubungkan ke database
    db = ConnectDB()

    // Definisikan rute
    userController := controllers.NewUserController(db)

    router.GET("/users", userController.GetAllUsers)
    router.POST("/users", userController.CreateUser)
    router.PUT("/users/:id", userController.UpdateUser)
    router.DELETE("/users/:id", userController.DeleteUser)

    // Jalankan server
    router.Run(":8081")
}

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
