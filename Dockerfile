FROM golang:1.21.0-alpine

WORKDIR /mini-project

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o main.app .

EXPOSE 8000

CMD ["/mini-project/main.app"]