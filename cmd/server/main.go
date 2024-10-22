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
	accountController := controllers.NewCreateAccountController(accountDB)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/", accountController.Create)
	//r.Get("/balance/{id}", accountController.Handle)

	err = http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}
