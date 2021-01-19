package shared

import (
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
)

func GetUser(tokenString string) (user string) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		hmacSampleSecret := []byte("bicycle")
		return hmacSampleSecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		//fmt.Println(claims["name"], claims["sub"])
		sub := claims["sub"]
		user := fmt.Sprintf("%v", sub)
		return user
	} else {
		fmt.Println(err)
	}
	return
}
