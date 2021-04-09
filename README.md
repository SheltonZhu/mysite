# CLI备忘录

此项目用于学习gin，grpc，gorm，docker，docker-compose，nginx 负载均衡（master，slave）

## Docker

构建app：

`docker build -t <Name:Tag> --target prod .`

构建nginx:

`docker build -t <Name:Tag> --target nginx .`

创建网络：

`docker network create <Network Name>`

# Docker-compose

启动：

`docker-compose --env-file=prod.env up -d`

多slave：

`docker-compose --env-file=prod.env up -d --scale slave=4`

多master：

`docker-compose --env-file=prod.env up -d --scale master=4`

# Protobuf

`protoc --go_out=. <Proto file path>`

# Grpc

`protoc --go-grpc_out=. <Proto file path>`