FROM golang:1.24-alpine

WORKDIR /app

COPY go.mod ./
RUN go mod tidy

COPY . .

RUN go build -o main .

ENV APP_ENV=development

CMD ["./main"]