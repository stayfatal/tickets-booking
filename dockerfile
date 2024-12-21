FROM golang:1.22-alpine3.20 AS builder

WORKDIR /tickets-booking

COPY go.mod go.sum ./
RUN go mod download

COPY . .

WORKDIR /tickets-booking/services/auth
RUN go build -o ./cmd/app/app ./cmd/app/main.go

WORKDIR /tickets-booking/services/gateway
RUN go build -o ./cmd/app/app ./cmd/app/main.go

FROM alpine:3.20 AS auth
WORKDIR /tickets-booking
COPY --from=builder /tickets-booking /tickets-booking
CMD ["/tickets-booking/services/auth/cmd/app/app"]

FROM alpine:3.20 AS gateway
WORKDIR /tickets-booking
COPY --from=builder /tickets-booking /tickets-booking
CMD ["/tickets-booking/services/gateway/cmd/app/app"]

FROM golang:1.22-alpine3.20 AS test
WORKDIR /tickets-booking
COPY . .
CMD ["go", "test", "-v", "./..."]
