package main

import (
	"fmt"
	"log"

	database "github.com/gipsh/stori-challenge/internal/db"
	"github.com/gipsh/stori-challenge/internal/mailer"
	"github.com/gipsh/stori-challenge/internal/parser"
	"github.com/gipsh/stori-challenge/internal/repository"
	"github.com/gipsh/stori-challenge/internal/service"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := database.Connection()
	defer database.Close()

	err = database.Migrate(db)
	if err != nil {
		panic(err)
	}

	mailer := mailer.NewMailer()
	repo := repository.NewRepository(db)
	svc := service.NewService(mailer, repo)
	parser := parser.NewParser()

	txs, err := parser.ParseFile("txns.csv")
	if err != nil {
		panic(err)
	}

	summary := svc.GenerateSummary(txs)
	fmt.Println(summary)

	err = svc.SendSummary(summary)
	if err != nil {
		panic(err)
	}
}
