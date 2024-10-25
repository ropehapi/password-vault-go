package main

import (
	"database/sql"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/go-sql-driver/mysql"
	//_ "github.com/mattn/go-sqlite3"
	"github.com/ropehapi/password-vault-go/internal/infrastructure/database"
	"github.com/ropehapi/password-vault-go/internal/infrastructure/web/controllers"
	"log"
	"net/http"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/password_vault?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	accountDB := database.NewAccountRepository(db)
	accountController := controllers.NewAccountController(accountDB)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/account", accountController.Create)
	r.Get("/account", accountController.GetAll)
	r.Get("/account/{name}", accountController.GetByName)
	r.Delete("/account/{id}", accountController.Delete)

	err = http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}
