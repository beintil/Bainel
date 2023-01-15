package user

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	NickName string             `bson:"nickName" json:"nickName"`
	Email    string             `bson:"email" json:"email"`
	Password string             `bson:"password" json:"password"`
}
