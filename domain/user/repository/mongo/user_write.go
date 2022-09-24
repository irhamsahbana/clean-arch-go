package mongo

import (
	"ca-boilerplate/domain"
	"ca-boilerplate/lib/logger"
	"context"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

func (repo *userRepository) Register(ctx context.Context, user *domain.User) (*domain.User, int, error) {
	_, err := repo.Collection.InsertOne(ctx, user)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return user, http.StatusOK, nil
}

func (repo *userRepository) InsertToken(ctx context.Context, userId, tokenId string) (*domain.User, int, error) {
	filter := bson.M{"uuid": userId}
	// update := bson.M{"$push": bson.M{"tokens": tokenId}}
	update := bson.M{
		"$set": bson.M{
			"tokens": bson.M{
				"$ifNull": bson.A{
					bson.M{"$concatArrays": bson.A{"$tokens", bson.A{tokenId}}},
					bson.A{tokenId},
				},
			},
		},
	}

	_, err := repo.Collection.UpdateOne(ctx, filter, bson.A{update})
	if err != nil {
		logger.Log(logrus.Fields{
			"error": err,
		}).Error("Failed to insert token")
		return nil, http.StatusInternalServerError, err
	}

	user, code, err := repo.FindUserBy(ctx, "uuid", userId, false)
	if err != nil {
		logger.Log(logrus.Fields{
			"error": err,
		}).Error("Failed to find user")
		return nil, code, err
	}

	user.Tokens = []string{tokenId}
	return user, http.StatusOK, nil
}

func (repo *userRepository) RemoveToken(ctx context.Context, userId, tokenId string) (*domain.User, int, error) {
	filter := bson.M{"uuid": userId}
	update := bson.M{"$pull": bson.M{"tokens": tokenId}}

	_, err := repo.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		logger.Log(logrus.Fields{
			"error": err,
		}).Error("Failed to remove token")
		return nil, http.StatusInternalServerError, err
	}

	user, code, err := repo.FindUserBy(ctx, "uuid", userId, false)
	if err != nil {
		logger.Log(logrus.Fields{
			"error": err,
		}).Error("Failed to find user")
		return nil, code, err
	}

	user.Tokens = []string{tokenId}

	return user, http.StatusOK, nil
}

func (repo *userRepository) UpdateUser(ctx context.Context, userId string, data *domain.User) (*domain.User, int, error) {
	filter := bson.M{"uuid": userId}
	update := bson.M{"$set": data}

	_, err := repo.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		logger.Log(logrus.Fields{
			"error": err,
		}).Error("Failed to update user")
		return nil, http.StatusInternalServerError, err
	}

	user, code, err := repo.FindUserBy(ctx, "uuid", userId, false)
	if err != nil {
		logger.Log(logrus.Fields{
			"error": err,
		}).Error("Failed to find user")
		return nil, code, err
	}

	return user, http.StatusOK, nil
}

func (repo *userRepository) DeleteUser(ctx context.Context, userId string) (*domain.User, int, error) {
	filter := bson.M{"uuid": userId}
	update := bson.M{
		"$set": bson.M{
			"deleted_at": time.Now().UnixMicro(),
		},
	}

	_, err := repo.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		logger.Log(logrus.Fields{
			"error": err,
		}).Error("Failed to delete user")
		return nil, http.StatusInternalServerError, err
	}

	user, code, err := repo.FindUserBy(ctx, "uuid", userId, true)
	if err != nil {
		logger.Log(logrus.Fields{
			"error": err,
		}).Error("Failed to find user")
		return nil, code, err
	}

	return user, http.StatusOK, nil
}
