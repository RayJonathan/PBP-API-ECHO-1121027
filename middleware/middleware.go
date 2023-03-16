package middleware

import (
	"github.com/labstack/echo/middleware"
)

var IsAutenticated = middleware.JWTWithConfig(middleware.JWTConfig{

	SigningKey: []byte("UUID:RAHASIA"),
})
