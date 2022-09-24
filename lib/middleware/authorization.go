package middleware

import (
	"ca-boilerplate/lib/http_response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserRole string

const (
	UserRole_USER UserRole = "user"
)

func Authorization(permittedRoles []UserRole) gin.HandlerFunc {
	return func(c *gin.Context) {
		for _, role := range permittedRoles {
			if role == UserRole(c.GetString("user_role")) {
				c.Next()
				return
			}
		}

		http_response.ReturnResponse(c, http.StatusForbidden, http.StatusText(http.StatusForbidden), nil)
		c.Abort()
	}
}
