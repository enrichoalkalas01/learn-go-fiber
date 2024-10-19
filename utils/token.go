package utils

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt"
)

var secretKey = []byte("secret-key-from-env-here")

type CreateTokenParams struct {
	ExpiredNumber *int
	ExpiredType   *string
	SourceFrom    *string
	TokenType     *string
	UserData      *map[string]interface{}
}

type ResultCreateToken struct {
	Token          *string
	ExpToken       *int64
	ExpTokenString *string
}

func CreateTokenJWT(params *CreateTokenParams) ResultCreateToken {
	// Params Validator & Configuration Default
	filterCreateTokenParams(params)

	// Create Date Expired (String Date, Millisecond Date)
	dateFuture := getFutureDate(futureDateParams{ExpiredType: Hours, Value: 1})

	// Set Claim Token
	claims := jwt.MapClaims{
		"exp":    *dateFuture.Milliseconds,
		"iat":    time.Now().Unix(),
		"source": *params.SourceFrom,
		"type":   *params.TokenType,
	}

	// Add UserData From Object If Exist
	if params.UserData != nil {
		for key, value := range *params.UserData {
			claims[key] = value
		}
	}

	// Create token with klaim
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token dengan secret key
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		log.Fatalf("Error generating token: %v", err)
	}

	formattedDateString := *dateFuture.FormattedDate
	millisecondsDateString := *dateFuture.Milliseconds

	// fmt.Println(tokenString)
	// fmt.Println(formattedDateString)
	// fmt.Println(*dateFuture.Milliseconds)
	// fmt.Println(*dateFuture.FutureTime)
	// fmt.Println(time.Now())
	// fmt.Println(time.UnixMilli(*dateFuture.Milliseconds))

	// Passing Results
	return ResultCreateToken{
		Token:          &tokenString,
		ExpToken:       &millisecondsDateString,
		ExpTokenString: &formattedDateString,
	}
}

// Fungsi filter untuk memastikan nilai default diterapkan
func filterCreateTokenParams(params *CreateTokenParams) {
	// Filter ExpiredNumber
	if params.ExpiredNumber == nil {
		defaultValue := 1
		params.ExpiredNumber = &defaultValue
	}

	// Filter ExpiredType
	if params.ExpiredType == nil {
		defaultValue := "day" // Bisa diganti dengan "minutes", "hours", dll.
		params.ExpiredType = &defaultValue
	}

	// Filter SourceFrom
	if params.SourceFrom == nil {
		defaultValue := "default" // Sumber token, seperti "documentation-be" atau "account-be"
		params.SourceFrom = &defaultValue
	}

	// Filter TokenType
	if params.TokenType == nil {
		defaultValue := "token" // Bisa berupa "access token" atau "refresh token"
		params.TokenType = &defaultValue
	}

	// Filter UserData
	if params.UserData == nil {
		defaultValue := map[string]interface{}{
			"id":       1,
			"username": "",
			"email":    "",
			"type":     "",
		}
		params.UserData = &defaultValue
	}
}
