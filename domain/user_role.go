package domain

import (
	"ca-boilerplate/dto"
	"context"
)

type UserRole struct {
	UUID      string `bson:"uuid"`
	Name      string `bson:"name"`
	Power     int    `bson:"power"`
	CreatedAt int64  `bson:"created_at"`
	UpdatedAt *int64 `bson:"updated_at,omitempty"`
	DeletedAt *int64 `bson:"deleted_at,omitempty"`
}

type UserRoleUsecaseContract interface {
	UpsertUserRole(ctx context.Context, data *dto.UserRoleUpsertrequest) (*dto.UserRoleModel, int, error)
	// FindUserRole(ctx context.Context, id string, withTrashed bool) (*UserRoleResponse, int, error)
	// DeleteUserRole(ctx context.Context, id string) (*UserRoleResponse, int, error)
}

type UserRoleRepositoryContract interface {
	FindUserRole(ctx context.Context, id string, withTrashed bool) (*UserRole, int, error)
	FindUserRoleBy(ctx context.Context, key string, val interface{}, withTrashed bool) (*UserRole, int, error)
	UpsertUserRole(ctx context.Context, userRole *UserRole) (*UserRole, int, error)
	DeleteUserRole(ctx context.Context, id string) (*UserRole, int, error)
}
