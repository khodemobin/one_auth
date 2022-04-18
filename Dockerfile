FROM golang:alpine

COPY ./src /app

WORKDIR /app

RUN go mod tidy

RUN go build main.go

EXPOSE 3000

CMD ["sh","-c","./main"]
