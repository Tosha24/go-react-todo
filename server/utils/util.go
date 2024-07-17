package utils

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET_KEY"));

func LoadMongoURIFromEnv() string {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    return os.Getenv("MONGODB_URI")
}

func GenerateJWTToken(id string) (string, error) {
	// create a new token
	token := jwt.New(jwt.SigningMethodHS256)
	
	// set the claims
	claims := token.Claims.(jwt.MapClaims)
	claims["userId"] = id
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	
	// sign the token
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func AuthenticateJWTToken(tokenString string) (string, error) {
	// parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return "", err
	}
	
	// check if the token is valid
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// get the userId from the claims in form of string and return it
		return claims["userId"].(string), nil
	}
	return "", err
}

func HashPassword(password string) string {
	// hash the password using bcrypt and then return it
	bcryptPassword, err := bcrypt.GenerateFromPassword([]byte(password), 11)
	if err != nil {
		log.Fatal(err)
	}
	return string(bcryptPassword);
}

func ComparePasswords(hashedPassword string, password string) bool {
	// compare the hashed password with the password and return the result
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}