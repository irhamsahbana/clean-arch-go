package authuser

import (
	"ca-boilerplate/bootstrap"
	"ca-boilerplate/domain"
	"ca-boilerplate/dto"
	"context"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindUser(id string) (*dto.UserResponse, int, error) {
	c := context.Background()
	ctx, cancel := context.WithTimeout(c, 10*time.Second)
	defer cancel()

	DB := bootstrap.App.Mongo.Database(bootstrap.App.Config.GetString("mongo.name"))
	collUsers := DB.Collection("users")
	collUserRoles := DB.Collection("user_roles")

	filter := bson.M{
		"$and": bson.A{
			bson.M{"uuid": id},
			bson.M{"deleted_at": nil},
		},
	}

	resultUser := collUsers.FindOne(ctx, filter)
	if resultUser.Err() != nil {
		if resultUser.Err().Error() == mongo.ErrNoDocuments.Error() {
			return nil, http.StatusNotFound, resultUser.Err()
		}
		return nil, http.StatusInternalServerError, resultUser.Err()
	}

	var user domain.User
	err := resultUser.Decode(&user)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	filter = bson.M{
		"uuid": user.RoleUUID,
	}

	resultRole := collUserRoles.FindOne(ctx, filter)
	if resultRole.Err() != nil {
		if resultUser.Err().Error() == mongo.ErrNoDocuments.Error() {
			return nil, http.StatusNotFound, resultUser.Err()
		}
		return nil, http.StatusInternalServerError, resultRole.Err()
	}

	var userRole domain.UserRole
	err = resultRole.Decode(&userRole)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	var resp dto.UserResponse
	toDTO(&resp, &user, &userRole)

	return &resp, http.StatusOK, nil
}

func toDTO(resp *dto.UserResponse, userData *domain.User, userRoleData *domain.UserRole) {
	resp.UUID = userData.UUID
	resp.Name = userData.Name
	resp.Role = userRoleData.Name
	respCreatedAt := time.UnixMicro(userData.CreatedAt).UTC()
	resp.CreatedAt = &respCreatedAt
	if userData.UpdatedAt != nil {
		respUpdatedAt := time.UnixMicro(*userData.UpdatedAt).UTC()
		resp.UpdatedAt = &respUpdatedAt
	}
	if userData.DeletedAt != nil {
		respDeletedAt := time.UnixMicro(*userData.DeletedAt).UTC()
		resp.DeletedAt = &respDeletedAt
	}
}
