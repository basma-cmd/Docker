FROM golang:1.22-alpine AS builder 

LABEL authors="ahtalbi, iaboudou, boutmana" \
      version="latest" \
      description="Dockerfile for the ascci art web"

WORKDIR /app

COPY . . 

RUN go build -o main main.go

FROM alpine:latest 

WORKDIR /dockeriz

COPY --from=builder /app/main /dockeriz
COPY --from=builder /app/statics /dockeriz/statics
COPY --from=builder /app/templates /dockeriz/templates

RUN apk add --no-cache bash

RUN mkdir app

EXPOSE 8080

CMD ["./main"]