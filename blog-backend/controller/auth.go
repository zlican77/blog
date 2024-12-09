package controller

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		s := sessions.Default(c)
		if value := s.Get("phoneNumber"); value == "15659589183" {
			c.Next()
			return
		}

		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "You are not allowed to.",
			"code":  401,
		})
		c.Abort()
		return
	}
}
