package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)



type Response struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Data interface{} `json:"data"`
}


func GenerateHash(password string)(string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)	
	return string(bytes), err
}

func CompareHashAndPassword(hash, password string)(bool){
	err := bcrypt.CompareHashAndPassword([]byte(hash),[]byte(password))
	return err == nil
}

func GenerateJWTToken(email string)(string, error){
	secretKey := []byte(os.Getenv("JWT_SECRET_KEY"))
	fmt.Println(fmt.Sprintln(secretKey))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp": time.Now().AddDate(0,0,7).Unix(),
	})
	fmt.Println(fmt.Sprintln(token))
	tokenString, err := token.SignedString(secretKey)
	fmt.Println(fmt.Sprintln(tokenString))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyJWTToken(tokenString string)(error){
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	if err != nil{
		return err
	}
	if !token.Valid{
		return fmt.Errorf("invalid token")
	}
	return nil
}