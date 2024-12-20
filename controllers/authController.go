package controllers

import (
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gofiber/fiber/v2"
	"github.com/razyneko/React-Go-JWT-Auth/database"
	"github.com/razyneko/React-Go-JWT-Auth/models"
	"golang.org/x/crypto/bcrypt"
)

const SecretKey = "secret"

func Register(c *fiber.Ctx) error {

	var data map[string]string

	if err:= c.BodyParser(&data); err != nil {
		return err
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]),14)

	user := models.User{
		Name: data["name"],
		Email: data["email"],
		Password: password,
	}

	database.DB.Create(&user)
	// Send a string response to the client
	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}
	
	var user models.User

	database.DB.Where("email =?", data["email"]).First(&user)

	if user.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message":"incorrect password",
		})
	}

	
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer: strconv.Itoa(int (user.Id)),
		// Manually Created a pointer to jwt.Time
		ExpiresAt: &jwt.Time{Time: time.Now().Add(time.Hour * 24)}, // for one day
	})

	token, err := claims.SignedString([]byte(SecretKey))

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message":"could not log in",
		})
	}

	// storing token in cookie
	cookie := fiber.Cookie{
		Name: "jwt",
		Value: token,
		Expires: time.Now().Add(time.Hour * 24),
		// to restrict frontend from accesing this cookie
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}
//retrieving user data from cookie 
func User(c *fiber.Ctx) error {
		cookie := c.Cookies("jwt")
		token , err := jwt.ParseWithClaims(cookie,&jwt.StandardClaims{}, func(token *jwt.Token)(interface{},error){
			return []byte(SecretKey), nil
		})

		if err != nil {
			c.Status(fiber.StatusUnauthorized)
			return c.JSON(fiber.Map{
				"message":"unauthenticated",
			})
		}

		claims := token.Claims.(*jwt.StandardClaims)
		var user models.User
		database.DB.Where("id = ?", claims.Issuer).First(&user)

		return c.JSON(user)
}

func LogOut(c *fiber.Ctx) error {
	// removing cookie
	// create a new cookie we set the expiry time in the past
	cookie := fiber.Cookie{
		Name: "jwt",
		Value: "",
		Expires: time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "success",
	})
}
