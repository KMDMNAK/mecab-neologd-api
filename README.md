# Mecab-API
This repository serve mecab api using Golang and Neologd Dictionary.

# Routes

## POST : /proper/count

Content-Type : application/json

Request Body
```
{
    text : string
}
```

___

## POST : /proper/extract

Content-Type : application/json

Request Body
```
{
    text : string
}
```

# Environment
- PORT : api port (default:3000).

# Build
You can use docker to deploy API server.
```
docker-compose up -d mecab-api
```