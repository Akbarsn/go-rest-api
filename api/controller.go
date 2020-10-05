package api

import (
	"encoding/json"
	"fmt"
	"go-rest-api/model"
	"go-rest-api/util"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gorilla/mux"

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

func (a *App) ReadAllUserHandler(w http.ResponseWriter, r *http.Request) {
	users := []model.User{}
	a.db.Find(&users)
	util.RespondJSON(w, 200, users)
}

func (a *App) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	var user model.User
	vars := mux.Vars(r)

	id := vars["id"]
	result := a.db.Where("id = ?", id).First(&user)
	if result.RowsAffected != 0 {
		err := a.db.Delete(&user).Error
		if err != nil {
			util.RespondError(w, 500, "Delete error")
		} else {
			util.RespondJSON(w, 200, map[string]string{"message": "User Deleted"})
		}
	} else {
		util.RespondError(w, 400, "ID Not found")
	}
}

func (a *App) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user model.User

	r.ParseMultipartForm(32 << 20)

	imageName, err := FileUpload(r)
	if err != nil {
		fmt.Println(err)
		util.RespondError(w, 500, "Upload image failed")
		return
	}

	user.Username = r.FormValue("username")
	user.Password = r.FormValue("password")
	user.FullName = r.FormValue("fullName")
	user.PhotoProfile = imageName

	fmt.Println(user)

	a.db.Create(&user)

	util.RespondJSON(w, 200, user)
}

func (a *App) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user model.User
	vars := mux.Vars(r)

	id := vars["id"]
	result := a.db.Where("id = ?", id).First(&user)
	if result.RowsAffected != 0 {
		r.ParseMultipartForm(32 << 20)

		imageName, err := FileUpload(r)
		if err != nil {
			fmt.Println(err)
			util.RespondError(w, 500, "Upload image failed")
			return
		}

		user.Username = r.FormValue("username")
		user.Password = r.FormValue("password")
		user.FullName = r.FormValue("fullName")
		user.PhotoProfile = imageName

		a.db.Save(&user)

		util.RespondJSON(w, 200, user)
	} else {
		util.RespondError(w, 400, "ID Not found")
	}
}

func FileUpload(r *http.Request) (string, error) {
	file, handler, err := r.FormFile("photoProfile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return "", err
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	tempFile, err := ioutil.TempFile("upload", "upload-*.png")
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	tempFile.Write(fileBytes)

	return tempFile.Name(), nil
}
