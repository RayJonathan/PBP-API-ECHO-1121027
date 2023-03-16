package controllers

import (
	"fmt"
	"net/http"
	"os/user"
	"time"

	"example.com/m/helpers"
	"example.com/m/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

type Claims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

func Admin() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Hi, you have access!")
	}
}

func generateToken(user *user.User, expirationTime time.Time, secret []byte) (string, time.Time, error) {

	claims := &Claims{
		Name: user.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", time.Now(), err
	}

	return tokenString, expirationTime, nil
}

func Logout(c echo.Context) error {
	cookie := &http.Cookie{
		Name:     "username",
		Value:    "",
		Expires:  time.Now(),
		Secure:   false,
		HttpOnly: true,
	}

	fmt.Println(cookie)
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, map[string]string{
		"Message  ": "SUCCESS LOGOUT",
	})
}

func CheckLogin(c echo.Context) error {

	username := c.FormValue("username")
	password := c.FormValue("password")

	res, err := models.CheckLogin(username, password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"messages": err.Error(),
		})
	}
	if !res {
		return echo.ErrUnauthorized

	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["level"] = "applications"
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte("UUID:RAHASIA"))
	fmt.Print(t)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"Messages ": err.Error(),
		})
	}
	sessionToken := t
	expiresAt := time.Now().Add(120 * time.Second)

	cookie := new(http.Cookie)
	cookie.Name = username
	cookie.Value = sessionToken
	cookie.Expires = expiresAt
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, map[string]string{
		"Message  ": "SUCCESS SET COOKIE",
	})
}

// func readCookie(c echo.Context) error {
// 	cookie, err := c.Cookie("username")
// 	if err != nil {
// 		return err
// 	}
// 	fmt.Println(cookie.Name)
// 	fmt.Println(cookie.Value)
// 	return c.String(http.StatusOK, "read a cookie")
// }

func GenerateHashPassword(c echo.Context) error {

	pass := c.Param("password")

	hash, _ := helpers.HashPassword(pass)

	return c.JSON(http.StatusOK, hash)
}

func StoreUsers(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	res, err := models.StoreUsers(username, password)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}
