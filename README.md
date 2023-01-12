# Crud GoLang

Projeto desenvolvido no curso de GoLang do canal HunCoding

- [Playlist - Meu Primeiro CRUD em GoLang](https://www.youtube.com/playlist?list=PLm-xZWCprwYQ3gyCxJ8TR1L2ZnUOPvOpr)

## Packages

- [Godotenv](https://github.com/joho/godotenv)
- [Gin Web Framework](https://github.com/gin-gonic/gin)

## Exemplos de request

- createUser

``` curl
curl --location --request POST 'http://localhost:8080/createUser' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "jean.jacques@teste.com",
    "password": "12345",
    "name": "Jean Jacques",
    "age": 23
}'
```
