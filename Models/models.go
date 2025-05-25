package models

import (
    "context"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
)

type User struct {
    ID          primitive.ObjectID `bson:"_id,omitempty"`
    SrNo        int                `bson:"srno"`         
    FirstName   string             `bson:"first_name"`
    LastName    string             `bson:"last_name"`
    Email       string             `bson:"email"`            
    Password    string             `bson:"password"`        
    Role        string             `bson:"role"`             
    Phone       string             `bson:"phone"`            
    Address     []Address          `bson:"address,omitempty"`
    CreatedAt   time.Time          `bson:"created_at"`
    UpdatedAt   time.Time          `bson:"updated_at"`
    IsVerified  bool               `bson:"is_verified"`      
}


// Add logic to get next srno before insert
func GetNextSrNo(db *mongo.Database) (int, error) {
    count, err := db.Collection("users").CountDocuments(context.TODO(), bson.M{})
    return int(count + 1), err
}
