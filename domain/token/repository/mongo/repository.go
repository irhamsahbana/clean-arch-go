package mongo

import (
	"ca-boilerplate/domain"
	custom "ca-boilerplate/lib/custom_type"

	"go.mongodb.org/mongo-driver/mongo"
)

type tokenRepository struct {
	DB         mongo.Database
	Collection mongo.Collection
	Type       custom.TokenableType
}

func NewTokenMongoRepository(DB mongo.Database, t custom.TokenableType) domain.TokenRepositoryContract {
	return &tokenRepository{
		DB:         DB,
		Collection: *DB.Collection("tokens"),
		Type:       t,
	}
}
