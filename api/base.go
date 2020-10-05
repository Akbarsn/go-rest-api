package api

import (
	"fmt"
	"go-rest-api/config"
	"go-rest-api/database"
	"go-rest-api/model"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type App struct {
	router *mux.Router
	db     *gorm.DB
}

func (a *App) InitAndServe(c *config.Configuration) {
	a.db = database.InstantiateDB(&c.Database)

	a.router = mux.NewRouter()

	model.MigrateUser(a.db)
	a.SetRoute()

	addr := fmt.Sprintf("%s:%d", c.Server.Host, c.Server.Port)

	fmt.Println("Listening to", addr)
	log.Fatal(http.ListenAndServe(addr, a.router))
}
