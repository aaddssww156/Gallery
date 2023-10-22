FROM golang:1.20-alpine

WORKDIR /app

# COPY go.mod go.sum ./
COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o main .
RUN chmod +x main

EXPOSE 3234
CMD ["./main"]
