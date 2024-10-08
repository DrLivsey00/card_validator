FROM golang:1.23-alpine as builder

WORKDIR /app

COPY . .

RUN go mod download


RUN CGO_ENABLED=0 GOOS=linux go build -o card_checker ./cmd/app

CMD ["/card_checker"]

EXPOSE 8080

CMD ["./card_checker"]