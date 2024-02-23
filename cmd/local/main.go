package main

import (
	"log"
	"os"

	database "github.com/gipsh/stori-challenge/internal/db"
	"github.com/gipsh/stori-challenge/internal/mailer"
	"github.com/gipsh/stori-challenge/internal/reader"
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
	//fileReader := reader.NewFileReader(reader.S3)
	fileReader := reader.NewFileReader(reader.Local)

	svc := service.NewService(mailer, repo, fileReader)

	txs, err := svc.ProcessFile(os.Getenv("PROCESS_FILE"))
	if err != nil {
		panic(err)
	}

	summary := svc.GenerateSummary(txs)
	log.Println(summary)

	err = svc.SendSummary(summary)
	if err != nil {
		panic(err)
	}
}
