FROM golang:1.25.3-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o sikabiz ./cmd/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app

COPY --from=builder /app/sikabiz .

COPY --from=builder /app/migrations ./migrations

COPY --from=builder /app/users_data.json .

EXPOSE 80

CMD ["./sikabiz", "server"]

