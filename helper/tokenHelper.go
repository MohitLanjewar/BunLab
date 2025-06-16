package helper

import (
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/mongo"
)

type SignedDetails struct {
	Email      string
	First_name string
	Last_name  string
	User_type  string
	Uid        string
	jwt.StandardClaims
}

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "users")

var SECREAT_KEY string = os.Getenv("SECREAT_KEY")


func GenerateAllTokens(email, firstName, lastName, userType, uid string) (signedToken string, signedRefreshToken string, err error) {
	
}
