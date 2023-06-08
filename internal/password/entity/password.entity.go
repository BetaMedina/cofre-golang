package user

import "go.mongodb.org/mongo-driver/bson/primitive"

type PasswordEntity struct {
	Id          primitive.ObjectID `json:"id" bson:"_id"`
	Name        string             `json:"email"`
	Description string             `json:"description"`
	Pass        string             `json:"password,omitempty"`
	HashedKey   string             `json:"-"`
}
