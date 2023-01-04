package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id       primitive.ObjectID `bson:"_id" json:"id"`
	Password string             `bson:"password" json:"password"`
	Email    string             `bson:"email" json:"email"`
	Roles    []string           `bson:"roles" json:"roles"`
} // @name User
