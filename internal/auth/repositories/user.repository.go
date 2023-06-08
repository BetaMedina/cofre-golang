package user

import (
	user "secrets-golang/internal/user/dto"
	entity "secrets-golang/internal/user/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
)

type UserRepository interface {
	Save(payload *user.UserPayloadDto) (*mongo.InsertOneResult, error)
	FindOne(email string) *entity.UserEntity
	FindById(id string) *entity.UserEntity
}

type userRepository struct {
	collection *mongo.Collection
}

func (u userRepository) Save(payload *user.UserPayloadDto) (*mongo.InsertOneResult, error) {
	result, err := u.collection.InsertOne(context.TODO(), bson.M{"nickName": payload.Nickname, "email": payload.Email, "password": payload.Password})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u userRepository) FindOne(email string) *entity.UserEntity {
	var output *entity.UserEntity
	err := u.collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&output)
	if err != nil {
		return nil
	}
	return output
}

func (u userRepository) FindById(id string) *entity.UserEntity {
	var output *entity.UserEntity
	objID, _ := primitive.ObjectIDFromHex(id)
	err := u.collection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&output)
	if err != nil {
		return nil
	}
	return output
}

func NewUserRepository(collection *mongo.Collection) UserRepository {
	return &userRepository{collection: collection}
}
