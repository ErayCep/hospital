package middleware

import (
	"hospital/db"
	"hospital/helpers"
	"hospital/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context) {
	// Get cookie from request
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	token, err := helpers.ParseToken(tokenString)

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Check the expire date
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// Find the user with ID in JWT token
		var staff models.Staff
		db.GetStaffWithID(&staff, claims["sub"].(float64))
		if staff.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// Attach to request
		c.Set("staff", staff)

		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}

func RequirePrivileged(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	token, err := helpers.ParseToken(tokenString)

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Check the expire date
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// Find the user with ID in JWT token
		var staff models.Staff
		db.GetStaffWithID(&staff, claims["sub"].(float64))
		if staff.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		if !staff.Privileged {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// Attach to request
		c.Set("staff", staff)

		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
