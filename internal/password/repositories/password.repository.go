package user

import (
	password "secrets-golang/internal/password/dto"
	entity "secrets-golang/internal/password/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
)

type PasswordRepository interface {
	Save(payload *password.SavePasswordRepositoryDto) (*mongo.InsertOneResult, error)
	List(userId string) (*[]password.ListPasswordRepositoryDto, error)
	Read(id string, userId string) (*entity.PasswordEntity, error)
}

type passwordRepository struct {
	collection *mongo.Collection
}

func (u passwordRepository) Save(payload *password.SavePasswordRepositoryDto) (*mongo.InsertOneResult, error) {
	result, err := u.collection.InsertOne(context.TODO(), bson.M{
		"name":        payload.Name,
		"description": payload.Description,
		"hashedKey":   payload.HashedKey,
		"userId":      payload.UserId,
		"pass":        payload.Pass,
		"platform":    payload.Platform,
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u passwordRepository) List(userId string) (*[]password.ListPasswordRepositoryDto, error) {
	var output []password.ListPasswordRepositoryDto
	current, err := u.collection.Find(context.TODO(), bson.M{
		"userId": userId,
	})
	if err != nil {
		return nil, err
	}
	for current.Next(context.TODO()) {
		var row password.ListPasswordRepositoryDto
		current.Decode(&row)
		output = append(output, row)
	}

	return &output, nil
}

func (u passwordRepository) Read(id string, userId string) (*entity.PasswordEntity, error) {
	var output *entity.PasswordEntity
	objectId, _ := primitive.ObjectIDFromHex(id)
	err := u.collection.FindOne(context.TODO(), bson.M{
		"_id":    objectId,
		"userId": userId,
	}).Decode(&output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func NewPasswordRepository(collection *mongo.Collection) PasswordRepository {
	return &passwordRepository{collection: collection}
}
