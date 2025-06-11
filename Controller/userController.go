package controller

import "go.mongodb.org/mongo-driver/mongo"

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "users")
validator := validator.New()

func HashPAssword()

func VerifyPassword()

func Signup()

func Login()

func GetUsers()

func GetUser()
