# api-golang-graphql

This is an API responding to GraphQL queries and mutations sent by an Angular-Apollo application.

It is fully explained [here](https://medium.com/@maumribeiro/my-first-graphql-server-in-go-responding-to-apollo-bd1c11426572)

## Pre-requisites

This project depends on [Gorilla mux](http://www.gorillatoolkit.org/pkg/mux) and [graphql-go](https://github.com/graphql-go/graphql) libraries.

Install them running

`$ go get github.com/gorilla/mux` and `$ go get github.com/graphql-go`

## To run it

After installing the pre-requisites, you can start the application browsing to its root folder and running `$ go run main.go`

## To test it

You can use [this application test](https://github.com/morris-ribs/graphql-apolloclient-angular) to test it.
