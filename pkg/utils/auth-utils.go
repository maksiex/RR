package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
	"time"
)

var key = []byte(os.Getenv("AUTH_KEY"))

func CreateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour).Unix()
	tokenStr, err := token.SignedString(key)

	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	return tokenStr, nil
}

func ValidateJWT(next func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header["Token"][0], func(t *jwt.Token) (interface{}, error) {
				_, ok := t.Method.(*jwt.SigningMethodHMAC)
				if !ok {
					w.WriteHeader(http.StatusUnauthorized)
					_, writeErr := w.Write([]byte("not authorized"))
					if writeErr != nil {
						fmt.Println("error while writing")
					}
				}
				return key, nil
			})

			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				_, writeErr := w.Write([]byte("not authorized: " + err.Error()))
				if writeErr != nil {
					fmt.Println("error while writing")
				}
			}

			if token.Valid {
				next(w, r)
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			_, writeErr := w.Write([]byte("not authorized"))
			if writeErr != nil {
				fmt.Println("error while writing")
			}
		}
	})
}

func GetJwt(w http.ResponseWriter, r *http.Request) {
	if r.Header["api"] != nil {
		if r.Header["api"][0] == "api_key" {
			token, err := CreateJWT()
			if err != nil {
				return
			}
			fmt.Println(w, token)
		}
	}
}
