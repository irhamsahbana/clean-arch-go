package mongo

import (
	"ca-boilerplate/domain"
	"context"
	"errors"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (repo *tokenRepository) FindTokenWithATandRT(ctx context.Context, accessToken, refreshToken string) (token *domain.Token, code int, err error) {
	filter := bson.D{
		{Key: "access_token", Value: accessToken},
		{Key: "refresh_token", Value: refreshToken},
	}

	doc := domain.Token{}

	err = repo.Collection.FindOne(ctx, filter).Decode(&doc)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, http.StatusNotFound, errors.New("token not found")
		}

		return nil, http.StatusInternalServerError, err
	}

	return &doc, http.StatusOK, nil
}
