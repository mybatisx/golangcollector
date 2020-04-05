#!/bin/sh

git pull origin master

docker build -t lanrenshipu:v1 .
docker rm -f lanrenshipu
docker run -d -p 10002:8080 --name lanrenshipu -v /home/assets:/home/assets lanrenshipu:v1

docker rm -f nginx
docker run \
  --name nginx \
  -d -p 81:80 \
  -v /home/github.com/nginx/html:/usr/share/nginx/html \
  -v /home/github.com/nginx/nginx.conf:/etc/nginx/nginx.conf:ro \
  -v /home/github.com/nginx/conf.d:/etc/nginx/conf.d \
  nginx