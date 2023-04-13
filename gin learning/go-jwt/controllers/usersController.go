package controllers

import (
	"go-jwt/initializers"
	"go-jwt/models"
	"net/http"

	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	//GET email/password for body request
	var body struct {
		Email    string
		Passeord string
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read  body",
		})
		return
	}
	//hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Passeord), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		return
	}
	//crete the user
	user := models.User{
		Email: body.Email, Password: string(hash)}

	result := initializers.DB.Create(&user) // pass pointer of data to Create

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})
		return
	}

	//respond
	c.JSON(http.StatusOK, gin.H{})
}
func Login(c *gin.Context) {
	//get the email and pass off req body
	var body struct {
		Email    string
		Passeord string
	}
	//}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read  body",
		})
		return

	}

	//look up request user
	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	//compare sent in pass with saved yser pass hash

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Passeord))
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})
		return
	}
	//send it back
	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}
