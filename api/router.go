package api

func (a *App) SetRoute() {
	a.router.HandleFunc("/ping", PingHandler).Methods("GET")
	a.router.HandleFunc("/login", a.LoginHandler).Methods("POST")
}
