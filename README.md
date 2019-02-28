# chi-http

Chi-http example for golang

[Chi](https://github.com/go-chi/chi) is great library for creating HTTP services and it has easy routing system.

## Motivation

When I started learning golang, it was little difficult to find basic examples.
This repository will try to address some problems that I encountered while learning golang.

This Repository will have basic examples to implement your APIs.

## Installation

For installing all the dependencies run **go get -d ./...**

This project didn't update with go mods.

## start server

You can start server by running command **go run main.go**

if you want to restart server on file change follow below steps

1. Install [realize](https://github.com/oxequa/realize) by running **go get github.com/oxequa/realize**

2. start server by **realize start**

## Table of Contents
* [main.go](https://github.com/saiumesh535/chi-http/blob/master/main.go)
main.go file which bootstrap our project

* [database/connectdb](https://github.com/saiumesh535/chi-http/blob/master/database/connectdb.go)
MongoDB connection example

* [Auth](https://github.com/saiumesh535/chi-http/tree/master/auth)
Folder contains login and sign up examples and also router Mount example.

* [download](https://github.com/saiumesh535/chi-http/blob/master/download/main.go)
Folder contains Downloading files from golang server example

* [static](https://github.com/saiumesh535/chi-http/blob/master/static/main.go)
Folder contains example for serving static file to browser such as html/css

* [upload](https://github.com/saiumesh535/chi-http/tree/master/upload) contains example for uploading files using golang

* [Private](https://github.com/saiumesh535/chi-http/tree/master/private)
Creating private routes for APIS, which requires JWT token

## some helpful resources for learning

* [GoBooks](https://github.com/dariubs/GoBooks)
* [Golangbot](https://golangbot.com/)
* [JustForFunc: Programming in Go](https://www.youtube.com/channel/UC_BzFbxG2za3bp5NRRRXJSw/featured)
* [Go by sentdex](https://www.youtube.com/watch?v=G3PvTWRIhZA)
* [Go reference book](https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/preface.html)

## Since working with JSON in golang is pretty confusing for many people, so making some of the good resources below.

*[JSON and Golang APIs](https://github.com/corylanou/tns-restful-json-api)
*[StackOverflow thread on JSON and Handling](https://stackoverflow.com/questions/15672556/handling-json-post-request-in-go)
