FROM golang:1.22.2-alpine3.19 AS builder
WORKDIR /goapiapp
COPY . /goapiapp/

RUN go build -o main main.go

FROM alpine:3.19
WORKDIR /goapiapp
COPY --from=builder /goapiapp/main .

EXPOSE 8080

CMD [ "/goapiapp/main" ]