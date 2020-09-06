# Swagger API Doc Guide

## Generate swagger.json

### Use swag (recommended)

#### 1. Install swag

```
go get -u github.com/swaggo/swag/cmd/swag
```

#### 2. Genertate json

In the project root folder, excute command:

```
make swagger
```

### Use go-swagger

#### 1. Install swagger cli tool

- Homebrew/Linuxbrew
    
    ```
    brew tap go-swagger/go-swagger
    brew install go-swagger
    ```
    
- [more info](https://goswagger.io/install.html) 

#### 2. Use command to generate

- if you don't use go module

    ```
    swagger generate spec -o ./swagger.json
    ```
    
- if you have used go module

    ```
    GO111MODULE=on go mod vendor
    GO111MODULE=off swagger generate spec -o ./swagger.json
    ```

## Visit API Doc

http://127.0.0.1/swagger/index.html
