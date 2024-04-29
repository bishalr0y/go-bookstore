FROM golang:1.22.2-alpine3.19

WORKDIR /home/app

EXPOSE 9010

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . . 

RUN go build -C cmd -o ../bin/go-bookstore

CMD [ "./bin/go-bookstore" ]