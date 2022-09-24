package http

import (
	"ca-boilerplate/domain"
	"ca-boilerplate/dto"
	"ca-boilerplate/lib/http_response"

	// "ca-boilerplate/lib/middleware"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserRoleHandler struct {
	userRoleUsecase domain.UserRoleUsecaseContract
}

func NewUserRoleHandler(router *gin.Engine, usecase domain.UserRoleUsecaseContract) {
	handler := &UserRoleHandler{
		userRoleUsecase: usecase,
	}

	// authorized := router.Group("/", middleware.Auth)
	router.PUT("/user-roles", handler.UpsertUserRole)
}

func (h *UserRoleHandler) UpsertUserRole(c *gin.Context) {
	var request dto.UserRoleUpsertrequest
	if err := c.BindJSON(&request); err != nil {
		http_response.ReturnResponse(c, http.StatusUnprocessableEntity, err.Error(), nil)
		return
	}

	ctx := context.Background()
	result, httpCode, err := h.userRoleUsecase.UpsertUserRole(ctx, &request)
	if err != nil {
		http_response.ReturnResponse(c, httpCode, err.Error(), nil)
		c.Abort()
		return
	}

	http_response.ReturnResponse(c, httpCode, "User Role Created", result)
}
