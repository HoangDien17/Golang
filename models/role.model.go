package models

import (
	"employee/app/enum"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Role struct {
	Id          primitive.ObjectID    `bson:"_id" json:"id"`
	Name        string                `bson:"name" json:"name"`
	Description string                `bson:"description" json:"description"`
	Permissions []enum.PermissionEnum `bson:"permissions" json:"permissions"`
} // @name Role
