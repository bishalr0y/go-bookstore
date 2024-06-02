FROM golang:1.22.2-alpine3.19 as builder

WORKDIR /home/app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . . 

RUN go build -C cmd -o ../bin/go-bookstore


# MULTISTAGE BUILD

# distro-less image
FROM gcr.io/distroless/static-debian12:latest 

COPY --from=builder /home/app/bin/go-bookstore .

EXPOSE 9010

CMD [ "./go-bookstore" ]