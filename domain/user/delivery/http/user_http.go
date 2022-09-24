package http

import (
	"ca-boilerplate/domain"
	"ca-boilerplate/dto"
	"ca-boilerplate/lib/http_response"
	jwthandler "ca-boilerplate/lib/jwt_handler"
	"ca-boilerplate/lib/middleware"
	"context"
	"net/http"

	"github.com/golang-jwt/jwt/v4"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserUsecase domain.UserUsecaseContract
}

func NewUserHandler(router *gin.Engine, usecase domain.UserUsecaseContract) {
	handler := &UserHandler{
		UserUsecase: usecase,
	}

	authorized := router.Group("/", middleware.Auth)
	authorized.GET("auth/logout", handler.Logout)

	router.POST("auth/login", handler.Login)
	router.POST("auth/register", handler.Register)
	router.GET("auth/refresh-token", handler.RefreshToken)

	router.POST("users", handler.Register)
	authorized.GET("users/:id", handler.Find)
	authorized.PUT("users/:id", handler.Update)
	authorized.DELETE("users/:id", handler.Delete)
}

func (h *UserHandler) Login(c *gin.Context) {
	var request dto.UserLoginRequest

	if err := c.BindJSON(&request); err != nil {
		http_response.ReturnResponse(c, http.StatusUnprocessableEntity, err.Error(), nil)
		return
	}

	ctx := context.Background()
	result, httpCode, err := h.UserUsecase.Login(ctx, &request)
	if err != nil {
		http_response.ReturnResponse(c, httpCode, err.Error(), nil)
		c.Abort()
		return
	}

	http_response.ReturnResponse(c, httpCode, "Authenticated", result)
}

func (h *UserHandler) Register(c *gin.Context) {
	var request dto.UserRegisterRequest

	if err := c.BindJSON(&request); err != nil {
		http_response.ReturnResponse(c, http.StatusUnprocessableEntity, err.Error(), nil)
		return
	}

	ctx := context.Background()
	result, httpCode, err := h.UserUsecase.Register(ctx, &request)
	if err != nil {
		http_response.ReturnResponse(c, httpCode, err.Error(), nil)
		c.Abort()
		return
	}

	http_response.ReturnResponse(c, httpCode, "Registered", result)
}

func (h *UserHandler) RefreshToken(c *gin.Context) {
	accessToken := c.GetHeader("X-ACCESS-TOKEN")
	refreshToken := c.GetHeader("X-REFRESH-TOKEN")

	_, err := jwthandler.ValidateToken(accessToken)
	if err != nil {
		v, _ := err.(*jwt.ValidationError)

		if v.Errors == jwt.ValidationErrorExpired {
		} else {
			http_response.ReturnResponse(c, http.StatusUnauthorized, err.Error(), nil)
			return
		}
	}

	claimsRT, err := jwthandler.ValidateToken(refreshToken)
	if err != nil {
		http_response.ReturnResponse(c, http.StatusUnauthorized, err.Error(), nil)
		return
	}

	ctx := context.Background()
	result, httpCode, err := h.UserUsecase.RefreshToken(ctx, accessToken, refreshToken, claimsRT.UserUUID)
	if err != nil {
		http_response.ReturnResponse(c, httpCode, err.Error(), nil)
		c.Abort()
		return
	}

	http_response.ReturnResponse(c, httpCode, "Token refreshed", result)
}

func (h *UserHandler) Logout(c *gin.Context) {
	AT := c.GetString("access_token")
	userId := c.GetString("user_uuid")

	ctx := context.Background()
	_, httpcode, err := h.UserUsecase.Logout(ctx, userId, AT)
	if err != nil {
		http_response.ReturnResponse(c, httpcode, err.Error(), nil)
		return
	}

	http_response.ReturnResponse(c, httpcode, "Logout", nil)
}

func (h *UserHandler) Find(c *gin.Context) {
	userId := c.Param("id")

	ctx := context.Background()
	result, httpCode, err := h.UserUsecase.FindUser(ctx, userId, false)
	if err != nil {
		http_response.ReturnResponse(c, httpCode, err.Error(), nil)
	}

	http_response.ReturnResponse(c, httpCode, "Found", result)
}

func (h *UserHandler) Delete(c *gin.Context) {
	userId := c.Param("id")

	ctx := context.Background()
	result, httpCode, err := h.UserUsecase.DeleteUser(ctx, userId)
	if err != nil {
		http_response.ReturnResponse(c, httpCode, err.Error(), nil)
	}

	http_response.ReturnResponse(c, httpCode, "Deleted", result)
}

func (h *UserHandler) Update(c *gin.Context) {
	var request dto.UserUpdateRequest

	userId := c.Param("id")

	if err := c.BindJSON(&request); err != nil {
		http_response.ReturnResponse(c, http.StatusUnprocessableEntity, err.Error(), nil)
		return
	}

	ctx := context.Background()
	result, httpCode, err := h.UserUsecase.UpdateUser(ctx, userId, &request)
	if err != nil {
		http_response.ReturnResponse(c, httpCode, err.Error(), nil)
		c.Abort()
		return
	}

	http_response.ReturnResponse(c, httpCode, "Updated", result)
}
