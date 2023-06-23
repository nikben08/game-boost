package middleware

import (
	"game-boost/config"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

var secretKey = []byte(config.Config("JWT_SECRET_KEY"))

func AuthMiddleware(c *fiber.Ctx) error {
	if c.Path() == "/api/v1/auth/signup" || c.Path() == "/api/v1/auth/signin" {
		return c.Next()
	}

	// Get the Authorization header value
	authHeader := c.Get("Authorization")

	// Check if the token exists
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	// Extract the token from the Authorization header
	tokenString := authHeader[7:] // Remove "Bearer " prefix

	// Parse and validate the JWT token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Make sure the signing method is HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.NewValidationError("Invalid signing method", jwt.ValidationErrorSignatureInvalid)
		}
		return secretKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Invalid token signature",
			})
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid token",
		})
	}

	// Check if the token is valid
	if !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid token",
		})
	}

	// Store the user ID from the token in the context
	claims := token.Claims.(jwt.MapClaims)
	c.Locals("userID", claims["userID"])

	// Call the next handler
	return c.Next()
}
