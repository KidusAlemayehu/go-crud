package views

import (
	"crud/go-crud/models"
	"crud/go-crud/utils"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func UsersList(ctx *gin.Context) {
	var users []models.User
	if err := utils.DB.Find(&users).Error; err != nil {
		ctx.AbortWithError(400, err)
		log.Fatal(err)
	}
	ctx.JSON(200, gin.H{
		"result": users,
	})
}

func UserList(ctx *gin.Context) {
	var user models.User
	id := ctx.Param("id")

	utils.DB.First(&user, id)
	ctx.JSON(200, gin.H{
		"result": user,
	})
}
func UserRemove(ctx *gin.Context) {
	var user models.User
	id := ctx.Param("id")

	if err := utils.DB.Unscoped().Delete(&user, id).Error; err != nil {
		ctx.AbortWithError(400, err)
		log.Fatal(err)
	}
	ctx.JSON(200, gin.H{
		"message": "Deleted successfully",
	})
}

func UserCreate(ctx *gin.Context) {
	var user models.User
	if !utils.DB.Migrator().HasTable(&user) {
		utils.DB.AutoMigrate(&user)
	}
	ctx.BindJSON(&user)
	time := time.Now()
	pwd, err := bcrypt.GenerateFromPassword([]byte(user.PASSWORD), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	password := string(pwd)
	// password = strings.ToValidUTF8(password, password)
	addUser := models.User{ID: user.ID, NAME: user.NAME, PHONE: user.PHONE, EMAIL: user.EMAIL, PASSWORD: password, CREATED_AT: time}
	if err := utils.DB.Create(&addUser).Error; err != nil {
		log.Fatal(err)
		ctx.AbortWithStatus(500)
		return
	}
	ctx.JSON(200, gin.H{
		"data": addUser,
	})
}

func UserLogin(ctx *gin.Context) {
	var input LoginInput
	var user models.User
	ctx.BindJSON(&input)

	if err := utils.DB.Where("name = ?", input.Username).First(&user).Error; err != nil {
		ctx.AbortWithError(400, err)
		log.Fatal(err)
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.PASSWORD), []byte(input.Password))
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"message": "Invalid password",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "user logged in successfully",
		"data":    user,
	})

}
