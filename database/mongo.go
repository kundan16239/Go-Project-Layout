package database

import (
	"context"
	"fmt"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectMongoDB establishes a connection to MongoDB and returns the database instance
func ConnectMongoDB() *mongo.Database {
	clientOptions := options.Client().ApplyURI(viper.GetString("mongodb_url"))
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		panic(err)
	}

	db := client.Database(viper.GetString("mongoDatabase"))

	fmt.Println("Connected to MongoDB!")

	return db
}
