# A Basic Implimentation of Blockchain Logic

- to run, clone to local the project

```
go mod tidy
go run main.go
```

this will initialise a blockchain with genisis block

- now place a POST request @ /new with JSON fields like

```
{
"title" : "Percy Jackson",
"author" : "Rick Riordan",
"isbn" : "1002",
"publish_date" : "2023-4-15"
}
```

- next create POST request @ / with id posted as book_id here

```
{
"book_id": "f294eb6a54068a5e2ee803b2e866668a",
"user":"Bjax",
"checkout_date" : "2023-7-11"
}
```

- lastly check out the chain with GET request @ /  this will return all created chains
