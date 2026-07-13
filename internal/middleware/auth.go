package middleware

import (
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(ctx *gin.Context) {
	header := ctx.GetHeader("Authorization")

	if header == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Authorization header required",
		})

		ctx.Abort()
		return
	}

	parts := strings.Split(header, " ")

	if len(parts) != 2 || parts[0] != "Bearer" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid authorization header",
		})
		ctx.Abort()
		return
	}

	tokenString := parts[1]

	// JWT verify karo
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		// Sirf HS256 allow karo
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil || !token.Valid {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid token",
		})
		ctx.Abort()
		return
	}

	// Claims nikalo
	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid claims",
		})
		ctx.Abort()
		return
	}

	// User ID nikalo
	userID := uint(claims["id"].(float64))

	// Context me save karo
	ctx.Set("userID", userID)

	// Next handler pe jao
	ctx.Next()
}
