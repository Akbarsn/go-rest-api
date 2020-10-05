package api

import (
	"net/http"
)

func (a *App) SetRoute() {
	a.router.HandleFunc("/ping", PingHandler).Methods("GET")
	a.router.HandleFunc("/login", a.LoginHandler).Methods("POST")

	a.router.Handle("/create", CheckToken(http.HandlerFunc(a.CreateUserHandler))).Methods("POST")
}
