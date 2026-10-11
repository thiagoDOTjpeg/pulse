FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o pulse_api ./cmd/...

FROM scratch

COPY --from=builder /app/pulse_api /pulse_api

EXPOSE 8000

ENTRYPOINT ["/pulse_api"]