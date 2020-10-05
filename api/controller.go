package api

import (
	"encoding/json"
	"go-rest-api/model"
	"go-rest-api/util"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GetSecretKey() []byte {
	return []byte("SecretKeyForGOlangAPI")
}

func PingHandler(w http.ResponseWriter, r *http.Request) {
	util.RespondJSON(w, 200, map[string]string{"message": "Pong"})
}

func (a *App) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var userDTO model.User
	var user model.User
	// var response map[string]string
	json.NewDecoder(r.Body).Decode(&userDTO)

	result := a.db.Where("username = ?", userDTO.Username).First(&user)
	if result.RowsAffected != 0 {
		if userDTO.Password == user.Password {
			expired := time.Now().Add(time.Hour * 24).Unix()
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"User":      user,
				"expiredAt": expired,
			})

			if tokenString, err := token.SignedString(GetSecretKey()); err != nil {
				util.RespondError(w, 500, "JWT Signing failed")
			} else {
				util.RespondJSON(w, 200, map[string]string{"token": tokenString})
			}
		} else {
			util.RespondError(w, 403, "Wrong password")
		}
	} else {
		util.RespondError(w, 403, "Username not found")
	}
}
