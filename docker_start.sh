#!/bin/sh

docker network create mysite
docker run -d --name master --network=mysite -e MYSQL="root:woaiyao@tcp(mysql:3306)/mysite?charset=utf8mb4&parseTime=True&loc=Local" mysite:v1.0.0 /root/app -rpc=true -rpc_port=8000 -http=false
docker run -d --name slave_1 --network=mysite mysite:v1.0.0 /root/app -master=master:8000
docker run -d --name slave_2 --network=mysite mysite:v1.0.0 /root/app -master=master:8000
docker run -d --name nginx -p 80:80 --network=mysite nginx

read