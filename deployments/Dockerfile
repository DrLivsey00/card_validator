FROM golang:1.23-alpine as builder

WORKDIR /app

COPY /go.mod /go.sum ./

RUN go mod download

COPY . .



COPY ../configs/.env /app/configs/.env

RUN CGO_ENABLED=0 GOOS=linux go build -o card_checker ./cmd/app

CMD ["/card_checker"]

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/card_checker .

COPY --from=builder /app/configs/.env /root/configs/.env

RUN mkdir -p /app/logs

EXPOSE 8080

CMD ["./card_checker"]