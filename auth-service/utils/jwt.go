package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var jwtKey = []byte("your-secret-key")

type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}


func GenerateToken(email string, duration time.Duration) (string, error) {
	// claims := jwt.MapClaims{
	// 	"email": email,
	// 	"exp":   time.Now().Add(duration).Unix(),
	// }

	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// return token.SignedString(jwtKey)
	claims := &Claims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate((time.Now().Add(duration))),
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}


	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ValidateToken(tokenStr string) (*Claims, error) {
	// token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
	// 	return jwtKey, nil
	// })
	// if err != nil || !token.Valid {
	// 	return "", errors.New("invalid token")
	// }

	// claims := token.Claims.(jwt.MapClaims)
	// email, ok := claims["email"].(string)
	// if !ok {
	// 	return "", errors.New("invalid email claim")
	// }

	// return email, nil
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{},error){
		//Vaildate the algorithm
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid{
		return nil, errors.New("invalid token")
	}

	return claims, nil
}


func GenerateRefreshToken() (string, error) {
	return uuid.New().String(), nil
}



// import (
// 	"errors"
// 	"time"

// 	"github.com/golang-jwt/jwt/v5"
// )

// var jwtSecret = []byte("secret-key") // TODO: Don't forget to create real secret key in .env


// func GenerateToken(userEmail string) (string, error) {
// 	claims := jwt.MapClaims{
// 		"email": userEmail,
// 		"exp": time.Now().Add(time.Hour * 72).Unix(), // 72 hour expiration for jwt token
// 	}
	
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

// 	return token.SignedString(jwtSecret)
// }

// // ValidateToken parses and validates a JWT token string
// func ValidateToken(tokenStr string) (jwt.MapClaims, error) {
// 	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
// 		// Check token method
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, errors.New("unexpected signing method")
// 		}
// 		return jwtSecret, nil
// 	})

// 	if err != nil {
// 		return nil, err
// 	}

// 	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
// 		return claims, nil
// 	}

// 	return nil, errors.New("invalid token")
// }