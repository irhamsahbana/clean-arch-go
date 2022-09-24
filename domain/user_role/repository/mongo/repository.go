package mongo

import (
	"ca-boilerplate/domain"

	"go.mongodb.org/mongo-driver/mongo"
)

type userRoleReposiotry struct {
	DB         mongo.Database
	Collection mongo.Collection
}

func NewUserRoleMongoRepository(DB mongo.Database) domain.UserRoleRepositoryContract {
	return &userRoleReposiotry{
		DB:         DB,
		Collection: *DB.Collection("user_roles"),
	}
}
