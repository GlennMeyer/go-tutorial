FROM golang:1.13.0

RUN go get -u github.com/gin-gonic/gin

WORKDIR /app

COPY . /app

RUN go build -o app

EXPOSE 8080

CMD "./app"
