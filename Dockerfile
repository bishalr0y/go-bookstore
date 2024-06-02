FROM golang:1.22.2-alpine3.19 as builder

WORKDIR /home/app

# EXPOSE 9010

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . . 

RUN go build -C cmd -o ../bin/go-bookstore

# CMD [ "./bin/go-bookstore" ]

# MULTISTAGE BUILD

FROM gcr.io/distroless/static-debian12:latest

COPY --from=builder /home/app/bin/go-bookstore .

CMD [ "./go-bookstore" ]