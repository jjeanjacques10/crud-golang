# Crud GoLang

Projeto desenvolvido no curso de GoLang do canal HunCoding

- [Playlist - Meu Primeiro CRUD em GoLang](https://www.youtube.com/playlist?list=PLm-xZWCprwYQ3gyCxJ8TR1L2ZnUOPvOpr)

## Packages

- [Godotenv](https://github.com/joho/godotenv)
- [Gin Web Framework](https://github.com/gin-gonic/gin)
- [GO Playground Validator](https://github.com/go-playground/validator)
- [Zap - Logging in Go](https://github.com/uber-go/zap)

## Environment

Arquivo .env

``` .env
LOG_OUTPUT=stdout
LOG_LEVEL=INFO
```

## Exemplos de request

- createUser

``` curl
curl --location --request POST 'http://localhost:8080/createUser' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "jean.jacques@teste.com",
    "password": "!23456*",
    "name": "Jean Jacques",
    "age": 23
}'
```
