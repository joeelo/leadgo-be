package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Program struct {
	Id           primitive.ObjectID `json:"id,omitempty"`
	Name         string             `json:"name,omitempty" validate:"required"`
	CoverImage   string             `json:"coverImage,omitempty" validate:"required"`
	Bio          string             `json:"bio,omitempty" validate:"required"`
	Href         string             `json:"href,omitempty" validate:"required"`
	PartnerUrl   string             `json:"partnerUrl,omitempty" validate:"required"`
	HelpsWith    []string
	Organization struct {
		program     bool
		summer      bool
		internship  bool
		scholarship bool
	}
}
