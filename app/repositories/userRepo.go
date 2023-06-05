package repositories

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository struct {
	Collection *mongo.Collection
	Db         *mongo.Database
}

func NewUserRepository(db *mongo.Database) *UserRepository {
	collection := db.Collection("user")
	return &UserRepository{
		Db:         db,
		Collection: collection,
	}
}

func (r *UserRepository) Create(ctx context.Context, doc interface{}) error {
	// Implement user creation logic using the MongoDB database connection
	_, err := r.Collection.InsertOne(ctx, doc)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetOne(ctx context.Context, doc interface{}) (map[string]interface{}, error) {
	// Implement user creation logic using the MongoDB database connection
	result := r.Collection.FindOne(ctx, doc)
	if result.Err() == mongo.ErrNoDocuments {
		return nil, nil
	}

	elem := make(bson.M)

	err := result.Decode(elem)
	if err != nil {
		return nil, err
	}

	return elem, nil
}

func (r *UserRepository) GetAll(ctx context.Context, doc interface{}) ([]map[string]interface{}, error) {
	// Implement user creation logic using the MongoDB database connection
	cursor, err := r.Collection.Find(ctx, doc)

	if err != nil {
		return nil, err
	}

	var elems []map[string]interface{}
	for cursor.Next(ctx) {
		// create a value into which the single document can be decoded
		var elem map[string]interface{}
		if err := cursor.Decode(&elem); err != nil {
			return nil, err
		}

		elems = append(elems, elem)
	}

	err = cursor.Close(ctx)
	if err != nil {
		return nil, err
	}
	return elems, nil

}

func (r *UserRepository) GetAllWithOptions(ctx context.Context, doc interface{}, options *options.FindOptions) ([]map[string]interface{}, error) {
	// Implement user creation logic using the MongoDB database connection
	cursor, err := r.Collection.Find(ctx, doc, options)

	if err != nil {
		return nil, err
	}

	var elems []map[string]interface{}
	for cursor.Next(ctx) {
		// create a value into which the single document can be decoded
		var elem map[string]interface{}
		if err := cursor.Decode(&elem); err != nil {
			return nil, err
		}

		elems = append(elems, elem)
	}

	err = cursor.Close(ctx)
	if err != nil {
		return nil, err
	}
	return elems, nil

}

func (r *UserRepository) FindOneAndUpdate(ctx context.Context, filter interface{}, update interface{}) (int64, error) {
	// Implement user creation logic using the MongoDB database connection
	result, err := r.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return 0, err
	}

	return result.ModifiedCount, nil

}

func (r *UserRepository) CountDocuments(ctx context.Context, doc interface{}) (int64, error) {
	// Implement user creation logic using the MongoDB database connection
	count, err := r.Collection.CountDocuments(ctx, doc)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *UserRepository) GetOneWithOptions(ctx context.Context, doc interface{}, options *options.FindOneOptions) (map[string]interface{}, error) {
	// Implement user creation logic using the MongoDB database connection
	result := r.Collection.FindOne(ctx, doc, options)
	if result.Err() == mongo.ErrNoDocuments {
		return nil, nil
	}

	elem := make(bson.M)

	err := result.Decode(elem)
	if err != nil {
		return nil, err
	}

	return elem, nil
}

func (r *UserRepository) Aggregation(ctx context.Context, pipeline interface{}) ([]map[string]interface{}, error) {
	cursor, err := r.Collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var results []map[string]interface{}
	if err := cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	return results, nil

}
