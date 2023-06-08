package user

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserEntity struct {
	Id       primitive.ObjectID `json:"_id" bson:"_id"`
	Email    string             `json:"email"`
	Nickname string             `json:"nickName"`
	Password string             `json:"password,omitempty"`
}
