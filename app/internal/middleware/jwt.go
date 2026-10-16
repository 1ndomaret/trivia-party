package middleware

import (
	"app/app/internal/helper"
	"os"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v5"
	"github.com/labstack/echo/v5"
)

func ParseJwtClaims(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c *echo.Context) error {
		token, err := echo.ContextGet[*jwt.Token](c, "user")
		if err != nil {
			return helper.Unauthorized(c, "Invalid or missing token")
		}

		claims, ok := token.Claims.(*helper.JwtCustomClaims)
		if !ok {
			return helper.Unauthorized(c, "Failed to parse claims")
		}

		c.Set("userID", claims.UserID)
		c.Set("name", claims.Name)
		c.Set("email", claims.Email)

		return next(c)
	}
}

func JwtConfig() echojwt.Config {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "secret"
	}

	return echojwt.Config{
		NewClaimsFunc: func(c *echo.Context) jwt.Claims {
			return new(helper.JwtCustomClaims)
		},
		ErrorHandler: func(c *echo.Context, err error) error {
			return helper.Unauthorized(c, "Missing or invalid token")
		},
		SigningKey: []byte(jwtSecret),
	}
}
