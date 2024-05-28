
### Important! before sending requests please get bearer token from /login endpoint !!

#### tech stack that i've used:
* docker for Containerization
* mysql for database
* go
* echo v4
* prometheus for monitoring

## make sure you've created an .env file according to .env.example !!!

```cp .env.example .env```


### To run docker dependencies please run the command below

```docker compose up```


### to run the application please run:
```go run cmd/main.go```


### you can use docker instead of the command above, the default port of application is 80. you can visit:
```http://<docker-machine-ip>:80```

