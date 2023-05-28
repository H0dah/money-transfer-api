# money-transfer-api
A system that emulates a RESTful API for money transfers between accounts.

## Dependencies
- GoLang go1.20.4

## Installation

First, start by cloning the repository:

```
git clone https://github.com/H0dah/money-transfer-api
```


- Once downloaded, access the project folder
```
cd money-transfer-api
```

- Install GoLang by follow this link and choose your system then follow instructions
```
https://go.dev/doc/install
```

## Run System

- Start the web application
```
go run money-transfer-api
```

## Run Tests

- while in project directory run this:
```
go test ./...
```

## Endpoints

- Endpoint to list all accounts in the system, request with GET method
```
http://localhost:8080/list
```


- Endpoint to make transfer between two accounts, request with POST method
```
http://localhost:8080/transfer
```

request body structure:
```
{
    "IdFrom": "id to transfer from",
    "Amount": float_number,
    "IdTo": "id to transfer to",
}
```

## Design Choices

- I choose not to use any package, simple is better
- Account package is the only who has access to accounts list ingested in the system, So no any other package can change it, it's completely separated
- Account package is separated from http so it is extandable
