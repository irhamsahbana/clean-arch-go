package mongo

import (
	"ca-boilerplate/domain"
	"context"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (repo *userRoleReposiotry) FindUserRole(ctx context.Context, id string, withTrashed bool) (*domain.UserRole, int, error) {
	var userRole domain.UserRole
	var filter bson.M

	if withTrashed {
		filter = bson.M{"uuid": id}
	} else {
		filter = bson.M{
			"$and": bson.A{
				bson.M{"uuid": id},
				bson.M{"deleted_at": bson.M{"$exists": false}},
			},
		}
	}

	err := repo.Collection.FindOne(ctx, filter).Decode(&userRole)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, http.StatusNotFound, err
		}

		return nil, http.StatusInternalServerError, err
	}

	return &userRole, http.StatusOK, nil
}

func (repo *userRoleReposiotry) FindUserRoleBy(ctx context.Context, key string, val interface{}, withTrashed bool) (*domain.UserRole, int, error) {
	var userRole domain.UserRole

	filter := bson.M{
		"$and": []bson.M{
			{key: val},
			{"deleted_at": bson.M{"$exists": false}},
		},
	}

	result := repo.Collection.FindOne(ctx, filter)
	err := result.Decode(&userRole)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return &userRole, http.StatusOK, nil
}
