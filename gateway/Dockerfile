FROM golang:1.21

WORKDIR /build

COPY ./gateway ./gateway
COPY ./shared ./shared
COPY ./hashing ./hashing

COPY go.work .
COPY go.work.sum .

RUN go work sync
RUN go build -o=gateway-service ./gateway/cmd/main.go

CMD [ "./gateway-service" ]