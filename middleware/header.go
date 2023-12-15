package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
)

var secretKey = []byte("eeSecretYouShouldHide")

func Header(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		reqToken := r.Header.Get("Authorization")
		splitToken := strings.Split(reqToken, "Bearer ")
		reqToken = splitToken[1]
		token, _ := jwt.Parse(reqToken, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", nil)
			}
			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return secretKey, nil
		})
		claims, _ := token.Claims.(jwt.MapClaims)
		fmt.Println(claims["user_id"])
		r = r.WithContext(context.WithValue(r.Context(), "user_id", claims["user_id"]))
		//.Set("user_id", claims["user_id"])
		next.ServeHTTP(w, r) /// every thing after this will be handle after trhe controller
	})
}
