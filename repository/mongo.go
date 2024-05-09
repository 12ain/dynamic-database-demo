package repository

import (
	"context"
	"dynamic-database-demo/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// NewMongoDBUserRepository creates a new MongoDBUserRepository instance
func NewMongoDBUserRepository(client *mongo.Client) UserRepository {
	return &MongoDBUserRepository{client: client}
}

// MongoDBUserRepository implements UserRepository interface using MongoDB as a backend
type MongoDBUserRepository struct {
	client *mongo.Client
}

func (m *MongoDBUserRepository) CreateUser(user model.User) error {
	collection := m.client.Database("dynamic-demo").Collection("users")
	_, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		return err
	}
	return nil
}

func (m *MongoDBUserRepository) GetUser(userID string) (model.User, error) {
	var user model.User
	collection := m.client.Database("dynamic-demo").Collection("users")
	filter := bson.M{"_id": userID}
	err := collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (m *MongoDBUserRepository) UpdateUser(user model.User) error {
	collection := m.client.Database("dynamic-demo").Collection("users")
	filter := bson.M{"_id": user.ID}
	update := bson.M{"$set": user}
	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (m *MongoDBUserRepository) DeleteUser(userID string) error {
	collection := m.client.Database("dynamic-demo").Collection("users")
	filter := bson.M{"_id": userID}
	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}
	return nil
}
