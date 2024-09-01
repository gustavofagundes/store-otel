# store-otel
project to implement opentelemtry partners of metrics and traces



# comandos docker
docker-compose -f .\assets\docker\docker-compose.yaml up -d
docker-compose -f .\assets\docker\docker-compose.yaml stop
docker-compose -f .\assets\docker\docker-compose.yaml rm

# Mysql
 mysql -u root -proot
 mysql -u user -ppass

# requests
## /items

this endpoints list all items present in the database

```bash
curl -v http://localhost:8080/items
```

##  /add

this endepoint you can add new itens in the database

```bash
curl --location 'http://localhost:8080/add' \
--header 'Content-Type: application/json' \
--data '{
    "items":[
    {
        "name":"iphone",
        "qtd":60,
        "price":860.99
    },
    {
        "name":"uno",
        "qtd":5,
        "price":700.00
    }

    ]
}'
```

## /buy

this endpoint You can select items present in the database and subtract the quantity specified on the list
```bash
curl --location 'http://localhost:8080/buy' \
--header 'Content-Type: application/json' \
--data '{
    "items":[
    {
        "name":"iphone",
        "qtd":5
    },
    {
        "name":"uno",
        "qtd":1
    }

    ]
}'
```