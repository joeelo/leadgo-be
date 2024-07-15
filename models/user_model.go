package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id            primitive.ObjectID `json:"id,omitempty"`
	Name          string             `json:"name,omitempty" validate:"required"`
	Email         string             `json:"email,omitempty" validate:"required"`
	Pronouns      string             `json:"pronouns,omitempty"`
	IsAdmin       bool               `json:"isAdmin,omitempty"`
	Grade         string             `json:"grade,omitempty" validate:"required"`
	SavedPrograms []primitive.ObjectID
}
