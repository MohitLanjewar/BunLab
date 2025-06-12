package controller

import "go.mongodb.org/mongo-driver/mongo"

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "users")
validator := validator.New()

func HashPAssword()

func VerifyPassword()

func Signup()

func Login()

func GetUsers()

func GetUser () gin.HandlerFunc{
	return func(c *gin.Context){
		userID:= c.Param("user_id")
	}
}
