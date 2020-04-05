#!/bin/sh

git pull origin master

docker build -t lanrenshipu:v1 .

docker run -d -p 10002:8080 -v /home/assets:/home/assets lanrenshipu:v1