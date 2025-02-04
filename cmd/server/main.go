package main

import (
	"database/sql"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"os"

	//_ "github.com/mattn/go-sqlite3"
	"github.com/ropehapi/password-vault-go/internal/infrastructure/database"
	"github.com/ropehapi/password-vault-go/internal/infrastructure/web/controllers"
	"log"
	"net/http"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:" + err.Error())
	}

	db, err := sql.Open(os.Getenv("DB_DRIVER"), fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME")))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	accountDB := database.NewAccountRepository(db)
	accountController := controllers.NewAccountController(accountDB)
	accountCodesDB := database.NewAccountCodesRepository(db)
	accountCodesController := controllers.NewAccountCodesController(accountCodesDB)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/account", accountController.Create)
	r.Get("/account", accountController.GetAll)
	r.Get("/account/{name}", accountController.GetByName)
	r.Delete("/account/{id}", accountController.Delete)
	r.Put("/account/{id}", accountController.Update)

	r.Post("/account-codes", accountCodesController.Create)
	r.Get("/account-codes", accountCodesController.GetAll)
	r.Get("/account-codes/{name}", accountCodesController.GetByName)
	r.Delete("/account-codes/{id}", accountCodesController.Delete)
	r.Put("/account-codes/{id}", accountCodesController.Update)

	err = http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}
