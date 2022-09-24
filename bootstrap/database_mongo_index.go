package bootstrap

import (
	"context"
	"fmt"
	"log"

	"github.com/fatih/color"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func initMongoDatabaseIndexes(ctx context.Context, client *mongo.Client, dbName string) {
	var collections []string = []string{
		"users",
		"user_roles",
		"tokens",

		"files",
	}

	// create collection
	// get all collections
	collNames, err := client.Database(dbName).ListCollectionNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	// create collection if not exists
	for _, collName := range collections {
		if !contains(collNames, collName) {
			err = client.Database(dbName).CreateCollection(ctx, collName)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	//delete all indexes first
	for _, collection := range collections {
		color.Yellow(fmt.Sprintf("deleting indexes from %s", collection) + " collection ...")
		_, err := client.Database(dbName).Collection(collection).Indexes().DropAll(ctx)
		if err != nil {
			color.Red("MongoDB: " + err.Error() + " on collection " + collection)
			log.Fatal(err)
		}
	}

	//create indexes
	for _, collection := range collections {
		switch collection {
		case "users":
			createCollectionIndex(collection)
			res, err := client.Database(dbName).Collection(collection).Indexes().CreateMany(ctx, []mongo.IndexModel{
				{
					Keys: bson.D{
						{Key: "uuid", Value: 1},
					},
					Options: options.Index().SetUnique(true),
				},
				{
					Keys:    bson.M{"email": 1},
					Options: options.Index().SetUnique(true),
				},
			})
			errCollectionIndexingCheck(err, collection)

			for _, index := range res {
				color.Green(fmt.Sprintf("index %s created", index))
			}

		case "user_roles":
			createCollectionIndex(collection)
			res, err := client.Database(dbName).Collection(collection).Indexes().CreateMany(ctx, []mongo.IndexModel{
				{
					Keys: bson.D{
						{Key: "uuid", Value: 1},
					},
					Options: options.Index().SetUnique(true),
				},
			})
			errCollectionIndexingCheck(err, collection)

			for _, index := range res {
				color.Green(fmt.Sprintf("index %s created", index))
			}

		case "tokens":
			createCollectionIndex(collection)
			res, err := client.Database(dbName).Collection(collection).Indexes().CreateMany(ctx, []mongo.IndexModel{
				{
					Keys: bson.D{
						{Key: "uuid", Value: 1},
					},
					Options: options.Index().SetUnique(true),
				},
				{
					Keys: bson.D{
						{Key: "access_token", Value: 1},
						{Key: "refresh_token", Value: 1},
					},
					Options: options.Index().SetUnique(true),
				},
			})
			errCollectionIndexingCheck(err, collection)

			for _, index := range res {
				color.Green(fmt.Sprintf("index %s created", index))
			}
		}
	}
}

func createCollectionIndex(collection string) {
	color.Cyan(fmt.Sprintf("creating indexes for %s", collection) + " collection ...")
}

func errCollectionIndexingCheck(err error, collection string) {
	if err != nil {
		color.Red("MongoDB: " + err.Error() + " on collection " + collection)
		log.Fatal(err)
	}
}
