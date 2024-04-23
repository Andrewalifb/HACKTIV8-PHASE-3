package config

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connection Database
func ConnectionDatabase(ctx context.Context, collName string) (*mongo.Collection, error) {
	mongoURI := os.Getenv("MONGO_DB_URI")
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		return nil, err
	}

	collection := client.Database("pt-creative-agency").Collection(collName)
	return collection, nil
}