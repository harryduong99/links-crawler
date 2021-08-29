package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Link struct {
	ID      primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"` // tag golang
	Url     string             `json:"url" bson:"url"`
	Crawled bool               `json:"crawled" bson:"crawled"`
	Domain  string             `json:"domain" bson:"domain"`
}
