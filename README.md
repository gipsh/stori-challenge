# Stori challenge 

Read a CSV file, generate the balance and send an email 

## Running with docker

For running the code locally use docker. 

First create a `.env` file based on the `example.env` provided 

Then build the image 

```bash 
docker build -t stori-challenge .
```

And the run the image 

```bash 
docker run -t stori-challenge
```

## Running locally

You can run the code directly on your machine 

first build the code 

```bash 
make build
```

then run it

```bash
RUN_LOCAL=true ./stori-challenge
```


## Running as lambda 

For the deploy i used [serverless framework](https://www.serverless.com/).

Everytime you upload an image to the bucket it trigers the lambda, the file 
is processed and the mail is sent.
All the definitions are located at `serverless.yml`


First compile the code 
```
make lambda
```

Then just deploy it, make sure you have your AWS enviroment already set

```
serverless deploy 
```

You should see something like this

![GitHub Image](/data/serverless-deploy.png)


## Mailer 

For local testing i used [mailtrap.io](https://mailtrap.io/) 
I implemented an SMTP, SES and a Dummy mailer.
You configure the mailer from env variables or from the `.env` file. 
The dummy mailer just print the body to standard output. 

Example of a generated summary mail: 

![example](/data/mailtrap-balance.png)

## File Reader 

There is an abstraction created so the user can read a file from local file system
or s3 bucket indistinctly.

## Database 

For simplicity i choosed [sqlite](https://www.sqlite.org/) as database engine.
It should be easy to extend the current code to support other engines. 
There is a migration to create the transactions table.

One drawback of using sqlite is that is not suitable for lambdas, as the filesystem 
can be deleted and you will end losing data. 

## Config 

The configuration is done using env variables. 
You can define those or use a `.env` file with all the configs
There is a `example.env` file with all the settings

The supported variables are: 

`RUN_LOCAL`  if true the code will run locally, using SMTP and local file system.
If not set it will run as a lambda using SES and fetching the file from s3 

`PROCESS_FILE` defines the name of the csv file to process. 

`SMTP_*` smtp config parameters

`FROM_EMAIL` email from value

`DB_DRIVER` currently the only supported is sqlite

`SQLITE_FILE` filename for storing the db

`S3_BUCKET`  Bucket name


## Test data

I generated a csv with 50 and other with 1000 transactions with this [tool](https://www.mockaroo.com/)
- [50txs](data/50txns.csv)
- [1000txs](data/1000txns.csv)





