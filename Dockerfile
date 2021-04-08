FROM golang:alpine as builder
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go env -w GO111MODULE=on
WORKDIR /go/src/mysite
COPY . .
RUN go get -v
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM  alpine:latest as prod
WORKDIR /root/
COPY --from=builder /go/src/mysite/conf/config.conf ./conf/
COPY --from=builder /go/src/mysite/app .
ENV GIN_MODE="release"
EXPOSE 8000 8000
CMD ["./app"]

FROM nginx:alpine as nginx
COPY ./nginx.conf /etc/nginx/conf.d/
RUN rm /etc/nginx/conf.d/default.conf