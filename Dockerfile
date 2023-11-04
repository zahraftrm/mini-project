FROM golang:1.21.0-alpine

WORKDIR /eduTrainerHub

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o main.app .

EXPOSE 8000

CMD ["/eduTrainerHub/main.app"]
