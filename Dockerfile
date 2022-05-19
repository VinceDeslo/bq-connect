FROM golang:1.18.1-alpine

WORKDIR /app

# Download necessary go
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy over all source files
COPY . .

RUN go build cmd/server/main.go

EXPOSE 8080

CMD ["./main"]