package mongo

import (
	"ca-boilerplate/domain"
	"ca-boilerplate/lib/logger"
	"context"
	"errors"
	"net/http"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (repo *userRepository) FindUserBy(ctx context.Context, key string, val interface{}, withTrashed bool) (*domain.User, int, error) {
	var user domain.User
	var filter bson.M

	if withTrashed {
		filter = bson.M{key: val}
	} else {
		filter = bson.M{
			"$and": bson.A{
				bson.M{key: val},
				bson.M{"deleted_at": nil},
			},
		}
	}

	err := repo.Collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			logger.Log(logrus.Fields{
				"key":   key,
				"value": val,
			}).Error("User not found")

			return nil, http.StatusNotFound, errors.New("user not found")
		}

		return nil, http.StatusInternalServerError, err
	}

	return &user, http.StatusOK, nil
}
