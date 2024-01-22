#!/bin/bash -x

git pull
go clean
go build
./snowman-booking
