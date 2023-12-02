
package middleware

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/opiziepazzle/golang-auth/models"
	"github.com/opiziepazzle/golang-auth/initializers"
)

func RequireAuth(c *fiber.Ctx) error {

		//Set some security headers:
	c.Set("X-XSS-Protection", "1; mode=block")
	c.Set("X-Content-Type-Options", "nosniff")
	c.Set("X-Download-Options", "noopen")
	c.Set("Strict-Transport-Security", "max-age=5184000")
	c.Set("X-Frame-Options", "SAMEORIGIN")
	c.Set("X-DNS-Prefetch-Control", "off")

	// Get the cookie from the request
	tokenString := c.Cookies("Authorization")
	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized: Missing token")
	}

	// Parse takes the token string and a function for looking up the key. The latter is especially
	 token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error){
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized: Invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized: Invalid token claims")
	}

	// check the exp
	exp, ok := claims["exp"].(float64)
	if !ok || float64(time.Now().Unix()) > exp {
		return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized: Token expired")
	}

	// Find the user with token sub
	var user models.User
	if err := initializers.DB.First(&user, claims["sub"].(string)).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized: User not found")
	}

	// Attach to req
	c.Locals("user", &user)

	fmt.Println(claims["foo"], claims["nbf"])

	// Go to next middleware
	return c.Next()
}







































