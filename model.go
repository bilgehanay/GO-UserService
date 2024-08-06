package main

import "go.mongodb.org/mongo-driver/bson/primitive"

type MongoConfig struct {
	ConnectionName   string `json:"connectionName"`
	ConnectionString string `json:"connectionString"`
	Collection       map[string]struct {
		N string `json:"n"` // name
		D string `json:"d"` // db
		C string `json:"c"` // col
	} `json:"collection"`
}

type User struct {
	ID       primitive.ObjectID `json:"id,omitempty," bson:"_id,omitempty"`
	Name     string             `json:"name,omitempty" bson:"name,omitempty" validate:"required,min=2,max=32"`
	Surname  string             `json:"surname,omitempty" bson:"surname,omitempty" validate:"required,min=2,max=32"`
	Email    string             `json:"email,omitempty" bson:"email,omitempty" validate:"required,email"`
	Password string             `json:"password,omitempty" bson:"password,omitempty" validate:"required,min=6,max=32"`
	Age      int                `json:"age,omitempty" bson:"age,omitempty" validate:"required,min=18,max=120"`
}
