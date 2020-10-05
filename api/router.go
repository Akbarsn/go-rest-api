package api

import (
	"net/http"
)

func (a *App) SetRoute() {
	a.router.HandleFunc("/ping", PingHandler).Methods("GET")
	a.router.HandleFunc("/login", a.LoginHandler).Methods("POST")

	a.router.Handle("/create", CheckToken(http.HandlerFunc(a.CreateUserHandler))).Methods("POST")
	a.router.Handle("/read", CheckToken(http.HandlerFunc(a.ReadAllUserHandler))).Methods("GET")
	a.router.Handle("/delete/{id}", CheckToken(http.HandlerFunc(a.DeleteUserHandler))).Methods("DELETE")
	a.router.Handle("/update/{id}", CheckToken(http.HandlerFunc(a.UpdateUserHandler))).Methods("PUT")
}
