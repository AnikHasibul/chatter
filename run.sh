#!/bin/bash
echo gopherjs build started
gopherjs build -m -o ./static/script.js ./app/main.go
echo go build started
go run main.go
