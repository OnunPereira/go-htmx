FROM golang:1.22.4-alpine3.19

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o /main

EXPOSE 8000

CMD ["/main"]
