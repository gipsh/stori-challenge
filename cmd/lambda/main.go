package main

import (
	"context"
	"database/sql"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
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

	ctx := context.Background()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// init aws config
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(os.Getenv("AWS_REGION")))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	db, err = database.Connection()
	if err != nil {
		log.Fatal(err)
	}

	err = database.Migrate(db)
	if err != nil {
		panic(err)
	}

	mailer := mailer.NewSESMailer(cfg, os.Getenv("FROM_EMAIL"))
	repo := repository.NewRepository(db)
	fileReader := reader.NewS3FileReader(cfg, os.Getenv("S3_BUCKET"))
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
		err := svc.ProcessFile(s3.Object.Key)
		if err != nil {
			log.Println(err)
			continue
		}
	}
}
