package password

import "go.mongodb.org/mongo-driver/bson/primitive"

type ListPasswordRepositoryDto struct {
	Id          primitive.ObjectID `json:"id" bson:"_id"`
	Description string             `json:"email,omitempty"`
	UserId      string             `json:"userId"`
	Name        string             `json:"name"`
	Platform    string             `json:"platform,omitempty"`
	Pass        string             `json:"-"`
}
