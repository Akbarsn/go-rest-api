package api

import (
	"context"
	"fmt"
	"go-rest-api/util"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

func CheckToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := strings.Split(r.Header.Get("Authorization"), "bearer ")
		if len(authHeader) != 2 {
			util.RespondError(w, 403, "Wrong token")
		} else {
			jwtToken := authHeader[1]
			token, _ := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}
				return GetSecretKey(), nil
			})
			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				user := claims["user"]
				ctx := context.WithValue(r.Context(), "user", user)
				r.WithContext(ctx)
				next.ServeHTTP(w, r)
				return
			} else {
				util.RespondError(w, 403, "Token not valid")
			}
		}
	})
}
