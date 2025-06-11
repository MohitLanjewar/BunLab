package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	SrNo         int                `bson:"srno"`
	FirstName    *string            `json:"first_name" validate:"required, min=2, max= 100"`
	LastName     *string            `json:"last_name " validate:"required, min=2, max= 100"`
	Email        *string            `json:"email" validate:"required,email"`
	Password     *string            `json:"password" validate:" required, min=6, max=100"`
	Phone        *string            `json:"phone" validate:"required"`
	Token        string             `json:"token"`
	RefreshToken *string            `bson:"refresh_token,omitempty"`
	ProfilePic   string             `bson:"profile_pic,omitempty"`
	Address      []Address          `bson:"address,omitempty"`
	CreatedAt    time.Time          `bson:"created_at"`
	UpdatedAt    time.Time          `bson:"updated_at"`
	IsVerified   bool               `bson:"is_verified"`
	RoleId       *string            `bson:"roleId"`
}

type Address struct {
	Primaryflag  bool   `bson:"primary_flag"`
	AddressType  string `bson:"address_type, omitempty"`
	AddressLine1 string `bson:"address_line1"`
	AddressLine2 string `bson:"address_line2,omitempty"`
	Pincode      string `bson:"pincode"`
	City         string `bson:"city"`
	State        string `bson:"state"`
	Country      string `bson:"country"`
}

// Add logic to get next srno before insert
func GetNextSrNo(db *mongo.Database) (int, error) {
	count, err := db.Collection("users").CountDocuments(context.TODO(), bson.M{})
	return int(count + 1), err
