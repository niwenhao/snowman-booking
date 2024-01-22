#!/bin/bash -x

cd /home/ubuntu/workspace/snowman-booking
git pull
go clean
go build
./snowman-booking
