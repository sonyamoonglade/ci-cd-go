package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"gihub.com/sonyamoonglade/ci-cd-go/internal/app"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {

	mux := http.NewServeMux()

	connstr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PWD"))

	db, err := sqlx.Open("postgres", connstr)
	if err != nil {
		log.Fatal(err.Error())
	}

	repo := app.NewRepo(db)
	service := app.NewService(repo)
	handler := app.NewHandler(service)

	handler.InitRoutes(mux)

	log.Fatal(http.ListenAndServe(":5000", mux))
}
