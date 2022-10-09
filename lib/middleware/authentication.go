package middleware

import (
	// authuser "ca-boilerplate/lib/auth_user"
	"ca-boilerplate/lib/http_response"
	jwthandler "ca-boilerplate/lib/jwt_handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	const BEARER_SCHEMA = "Bearer "

	authHeader := c.GetHeader("Authorization")

	if len(authHeader) < len(BEARER_SCHEMA) {
		http_response.ReturnResponse(c, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), nil)
		c.Abort()
		return
	}

	if authHeader == "" || authHeader[:len(BEARER_SCHEMA)] != BEARER_SCHEMA {
		http_response.ReturnResponse(c, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), nil)
		c.Abort()
		return
	}
	tokenString := authHeader[len(BEARER_SCHEMA):]

	_, err := jwthandler.ValidateToken(tokenString)
	if err != nil {
		http_response.ReturnResponse(c, http.StatusUnauthorized, err.Error(), nil)
		c.Abort()
		return
	}

	// user, code, err := authuser.FindUser(claims.UserUUID)
	// if err != nil {
	// 	http_response.ReturnResponse(c, code, err.Error(), nil)
	// 	c.Abort()
	// 	return
	// }

	// c.Set("access_token", tokenString)
	// c.Set("user_uuid", user.UUID)
	// c.Set("user_role", user.Role)

	c.Next()
}
