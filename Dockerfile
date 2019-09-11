FROM golang:1.13.0

RUN go get -u github.com/gin-gonic/gin github.com/go-pg/pg

COPY . .

RUN go build -o myapp .

EXPOSE 8080

CMD "./myapp"
