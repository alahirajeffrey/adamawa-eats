package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
 Id    primitive.ObjectID   `bson:"_id,omitempty"`
 Firstname  string      `bson:"firstName"`
 Lastname  string      `bson:"lastName"`
 Email    string             `bson:"email"`
 Password string             `bson:"password"`
 UserType string             `bson:"userType"`
}