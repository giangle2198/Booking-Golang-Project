package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `bson:"title,omitempty"`
	Description string             `bson:"description,omitempty"`
	Img         string             `bson:"img,omitempty"`
	Categories  []string           `bson:"categories,omitempty"`
	Sizes       []string           `bson:"sizes,omitempty"`
	Color       string             `bson:"color,omitempty"`
	Price       string             `bson:"price,omitempty"`
}
