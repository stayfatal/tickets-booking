FROM golang:1.22-alpine3.20 AS builder

WORKDIR /tickets-booking

COPY go.mod go.sum ./
RUN go mod download

COPY ./gen ./gen

COPY ./libs ./libs

FROM builder AS auth_builder

WORKDIR /tickets-booking

COPY ./services/auth ./services/auth

RUN go build -o app ./services/auth/cmd/app/main.go

FROM alpine:3.20 AS auth

WORKDIR /tickets-booking

COPY --from=auth_builder /tickets-booking /tickets-booking

CMD ["./app"]

FROM builder AS gateway_builder

WORKDIR /tickets-booking

COPY ./services/gateway ./services/gateway

RUN go build -o app ./services/gateway/cmd/app/main.go

FROM alpine:3.20 AS gateway

WORKDIR /tickets-booking

COPY --from=gateway_builder /tickets-booking /tickets-booking

CMD ["./app"]

FROM builder AS test

WORKDIR /tickets-booking

COPY ./services/auth ./services/auth

COPY ./services/gateway ./services/gateway

CMD ["go", "test", "-v", "./..."]
