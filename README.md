# CLI备忘录

## Docker

构建app：

`docker build -t <Name:Tag> --target prod .`

构建nginx:

`docker build -t <Name:Tag> --target nginx .`

创建网络：

`docker network create <Network Name>`

# Docker-compose

启动：

`docker-compose up -d`

# Protoc

`protoc --go_out=. <Proto file path>`