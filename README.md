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
The dummy mailer just print the body to standar output. 
It is possible to create also a SES mailer and use it for lambda deploy.

Example of a generated summary mail: 

![example](/data/mailtrap-balance.png)



## Database 

For simplicity i choose [sqlite](https://www.sqlite.org/) as database engine.
It should be easy to extend the current code to support other engines. 
There is a migration to create the transactions table.

## Config 

The configuration is done using env variables. 
You can define those or use a `.env` file with all the configs
There is a `example.env` file with all the settings








