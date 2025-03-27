package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// Tạo token JWT
func GenerateJWT(userID string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(), // Hết hạn sau 72 giờ
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
			// Get the Authorization header
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
					return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Missing token"})
			}

			// Check if the format is "Bearer <token>"
			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
					return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid token format"})
			}

			// Extract the token
			tokenString := parts[1]

			// Parse and validate the token
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
					// Validate the alg
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
							return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
					}
					
					// Return the secret key for validation
					return []byte(os.Getenv("JWT_SECRET")), nil
			})

			if err != nil || !token.Valid {
					return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid token"})
			}

			// Extract claims
			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
					return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Failed to extract claims"})
			}

			// Set user ID in context
			userID, ok := claims["user_id"].(string)
			if !ok {
					return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid user claim"})
			}
			
			c.Set("user_id", userID)
			
			return next(c)
	}
}