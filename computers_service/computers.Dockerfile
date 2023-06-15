FROM golang:1.19-alpine

WORKDIR /app/computers_service
COPY . .

RUN go mod download

CMD ["go", "run", "cmd/computers_service/main.go"]