# Stori challenge 

Read a CSV file, generate the balance and send an email 

## Running local

For running the code locally use docker. 

First create a `.env` file based on the `example.env` provided 

Then build the image 

```bash 
docker build -t stori-challege .
```

And the run the image 

```bash 
docker run -t stori-challenge
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

Then just deploy it

```
serverless deploy 
```

You should see something like this

![GitHub Image](/data/serverless-deploy.png)


## Mailer 

For local testing i used [mailtrap.io](https://mailtrap.io/) 
I implemented an STMP and a Dummy mailer.
You configure the mailer from env variables or from the `.env` file. 
The variable called `MAILER_METHOD` defines which implementation to use,
the valid values are `stmp` or `dummy`.  
The dummy mailer just print the body to standar output. 
It is possible to create also a SES mailer and use it for lambda deploy.

Example of a generated summary mail: 

![example](/data/mailtrap-balance.png)

## File Reader 

There is an abstraction created so the user can read a file from local file system
or s3 bucket indistinctly.

## Database 

For simplicity i choose [sqlite](https://www.sqlite.org/) as database engine.
It should be easy to extend the current code to support other engines. 
There is a migration to create the transactions table.

One drawback of using sqlite is that is not suitable for lambdas, as the filesystem 
can be deleted and you end losing data. 

## Config 

The configuration is done using env variables. 
You can define those or use a `.env` file with all the configs
There is a `example.env` file with all the settings

The supported variables are: 

`PROCESS_FILE` defines the name of the csv file to process. 

`MAILER_METHOD` valid values are smtp/dummy

`SMTP_*` smtp config parameters

`DB_DRIVER` currently the only supported is sqlite3

`SQLITE_FILE` filename for storing the db

`S3_BUCKET`  Bucket name







