FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .
RUN chmod +x main

# EXPOSE 8088
CMD ["./main"]
