package utils

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	return client.Database("gc1p3").Collection(collectionName)
}

// CheckDocumentExists checks if a document exists in the specified collection by its _id
func CheckDocumentExists(ctx context.Context, collection *mongo.Collection, id primitive.ObjectID) (bool, error) {
	count, err := collection.CountDocuments(ctx, bson.M{"_id": id}, options.Count().SetLimit(1))
	if err != nil {
		return false, err
	}
	return count > 0, nil
}


// UpdateOrdersStatus updates the status of all orders matching a specific status
func UpdateOrdersStatus(ctx context.Context, coll *mongo.Collection, currentStatus, newStatus string) (int64, error) {
	// Filter for orders with the current status
	filter := bson.M{
		"status": currentStatus,
	}

	// Update the status to the new status and update the updated_at field
	update := bson.M{
		"$set": bson.M{
			"status":     newStatus,
			"updated_at": time.Now(),
		},
	}

	// Perform the update
	result, err := coll.UpdateMany(ctx, filter, update)
	if err != nil {
		return 0, err
	}

	// Return the count of modified documents
	return result.ModifiedCount, nil
}
