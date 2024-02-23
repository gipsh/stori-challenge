package main

import (
	"context"
	"database/sql"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	database "github.com/gipsh/stori-challenge/internal/db"
	"github.com/gipsh/stori-challenge/internal/mailer"
	"github.com/gipsh/stori-challenge/internal/reader"

	"github.com/gipsh/stori-challenge/internal/repository"
	"github.com/gipsh/stori-challenge/internal/service"

	"github.com/joho/godotenv"
)

var (
	db  *sql.DB
	svc service.Service
)

// lambda entry point
func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err = database.Connection()
	if err != nil {
		log.Fatal(err)
	}

	err = database.Migrate(db)
	if err != nil {
		panic(err)
	}

	mailer := mailer.NewMailer()
	repo := repository.NewRepository(db)
	fileReader := reader.NewFileReader(reader.S3)
	svc = service.NewService(mailer, repo, fileReader)

	lambda.Start(handler)
}

// trigger by s3 event
func handler(ctx context.Context, s3Event events.S3Event) {

	for _, record := range s3Event.Records {
		s3 := record.S3
		log.Printf("[%s - %s] Bucket = %s, Key = %s \n",
			record.EventSource,
			record.EventTime,
			s3.Bucket.Name,
			s3.Object.Key)

		// process file
		txs, err := svc.ProcessFile(s3.Object.Key)
		if err != nil {
			log.Println(err)
			continue
		}

		// generate summary
		summary := svc.GenerateSummary(txs)

		// send summary
		err = svc.SendSummary(summary)
		if err != nil {
			log.Println(err)
			continue
		}
	}
}
