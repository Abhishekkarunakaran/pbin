FROM golang:1.23 AS builder

WORKDIR /app

COPY . .

RUN go mod download && go mod verify

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o pbin .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/pbin/ .
COPY src/view/static/ /app/src/view/static/
EXPOSE ${APP_PORT}

CMD ["./pbin"]
