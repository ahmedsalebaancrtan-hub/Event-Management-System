package middleware

import (
	"log/slog"
	"net/http"
	"strings"

	"github.com/ahmedsaleban/eventManagementsystem/infra"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type Claims struct {
	Sub  string `json:"sub"`
	Role string `json:"role"`
	jwt.StandardClaims
}

// Middleware-ka Access Token
func Authenticated() gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")

		// 1. Check header exists
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message":    "Missing Authorization header",
				"is_success": false,
			})
			return
		}

		// 2. Check Bearer format
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message":    "Invalid Authentication header format",
				"is_success": false,
			})
			return
		}

		// 3. Extract token
		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		secret := []byte(infra.Configuration.Access_jwt_Token)

		// 4. Parse token
		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			return secret, nil
		})

		// 5. Validate token
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message":    "Unauthenticated",
				"is_success": false,
			})
			return
		}

		// 6. Extract claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message":    "Invalid token claims",
				"is_success": false,
			})
			return
		}

		email := claims["sub"]
		role := claims["role"]

		// 7. Set context
		c.Set("email", email)
		c.Set("role", role)

		// 8. Log user
		slog.Info("Logged in User", "email", email)

		// 9. Continue
		c.Next()
	}
}
