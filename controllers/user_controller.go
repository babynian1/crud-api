
package controllers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "crud-api/models"
)

type UserController struct {
    db *gorm.DB
}

func NewUserController(db *gorm.DB) *UserController {
    return &UserController{db}
}

func (c *UserController) GetAllUsers(ctx *gin.Context) {
    var users []models.User
    c.db.Find(&users)
    ctx.JSON(http.StatusOK, users)
}

func (c *UserController) CreateUser(ctx *gin.Context) {
    var user models.User
    ctx.BindJSON(&user)
    c.db.Create(&user)
    ctx.JSON(http.StatusCreated, user)
}

func (c *UserController) UpdateUser(ctx *gin.Context) {
    var user models.User
    userID := ctx.Param("id")

    if err := c.db.First(&user, userID).Error; err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": "User tidak ditemukan"})
        return
    }

    ctx.BindJSON(&user)
    c.db.Save(&user)
    ctx.JSON(http.StatusOK, user)
}

func (c *UserController) DeleteUser(ctx *gin.Context) {
    var user models.User
    userID := ctx.Param("id")

    if err := c.db.First(&user, userID).Error; err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": "User tidak ditemukan"})
        return
    }

    c.db.Delete(&user)
    ctx.JSON(http.StatusOK, gin.H{"message": "User berhasil dihapus"})
}