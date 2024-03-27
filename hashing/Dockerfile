FROM golang:1.21

WORKDIR /build

COPY ./gateway ./gateway
COPY ./shared ./shared
COPY ./hashing ./hashing

COPY go.work .
COPY go.work.sum .

RUN go work sync
RUN go build -o=hashing-service ./hashing/cmd/main.go

CMD [ "./hashing-service" ]