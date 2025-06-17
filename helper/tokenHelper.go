package helper

import (
	"log" // Added log import for error handling
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/mongo"
)

type SignedDetails struct {
	Email      string
	First_name string
	Last_name  string
	User_type  string
	Uid        string
	jwt.RegisteredClaims
}

// Ensure userCollection and SECREAT_KEY are initialized correctly.
// For SECREAT_KEY, it's better to check if it's set and handle the case where it's not.
var userCollection *mongo.Collection = database.OpenCollection(database.Client, "users")

var SECRET_KEY string = os.Getenv("SECRET_KEY") // Corrected typo: SECREAT_KEY -> SECRET_KEY

func init() {
	if SECRET_KEY == "" {
		log.Fatal("SECRET_KEY environment variable not set")
	}
}

func GenerateAllTokens(email, firstName, lastName, userType, uid string) (signedToken string, signedRefreshToken string, err error) {
	// Access Token Claims (short-lived)
	claims := &SignedDetails{
		Email:      email,
		First_name: firstName,
		Last_name:  lastName,
		User_type:  userType,
		Uid:        uid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)), // Access token valid for 24 hours
			Issuer:    "BunLab",
		},
	}

	// Refresh Token Claims (long-lived, typically doesn't contain user-specific details beyond UID)
	// You might want to include Uid in refresh token claims as well if you need to identify the user
	// when refreshing the token without re-authenticating.
	refreshTokenClaims := &SignedDetails{
		Uid: uid, // Only include necessary details for refresh token, typically just the user ID
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 168)), // Refresh token valid for 7 days (168 hours)
			Issuer:    "BunLab",
		},
	}

	// Generate access token
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Printf("Error generating access token: %v", err) // Use Printf for more informative logging
		return "", "", err                                   // Return empty strings and the error
	}

	// Generate refresh token
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Printf("Error generating refresh token: %v", err)
		return "", "", err
	}

	return token, refreshToken, nil // Return nil for error if everything was successful
}

// UpdateToken updates the refresh token and sets a new access token
// This function is commonly used when a refresh token is used to get a new access token.
func UpdateToken(uid string, signedToken, signedRefreshToken string) error {
	// In a real application, you would typically store the refresh token (or its hash)
	// in your database associated with the user. When a user requests a new access token
	// using their refresh token, you would:
	// 1. Validate the refresh token.
	// 2. Look up the user in the database using the UID from the refresh token.
	// 3. Compare the provided refresh token with the one stored in the database.
	// 4. If valid, generate a new access token and a new refresh token (optional, but good for security).
	// 5. Update the stored refresh token in the database.

	// For simplicity, this example just demonstrates a placeholder for where you'd
	// interact with your database to update the tokens.
	// In a real scenario, you'd use the `userCollection` to query and update.

	// Example (conceptual):
	// _, err := userCollection.UpdateOne(
	// 	context.TODO(),
	// 	bson.M{"uid": uid},
	// 	bson.D{
	// 		{"$set", bson.M{
	// 			"token":         signedToken,
	// 			"refresh_token": signedRefreshToken,
	// 			"updated_at":    time.Now(),
	// 		}},
	// 	},
	// )
	// if err != nil {
	// 	return err
	// }
	return nil
}
