package utils

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func InitContext() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	return ctx, cancel
}

func GetErrorMessage(err error) string {
	if err != nil {
		return err.Error()
	}
	return ""
}

func GenerateToken(object *jwt.MapClaims) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, object)
	string, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))

	if err != nil {
		log.Println(err)
		return ""
	}

	return string
}

func RoleContains(all []string, s []string) bool {
	for _, v := range s {
		for _, v2 := range all {
			if v == v2 {
				return true
			}
		}
	}
	return false
}
