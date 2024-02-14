FROM golang:1.21-alpine as builder
RUN apk add --update git
RUN apk add --update gcc
RUN apk add --update g++
RUN apk add --update openssh-client
RUN apk add --update make

RUN mkdir /app
WORKDIR /app
ADD . .

RUN go get
RUN go build -o ./build/eth_pokhar


FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /
COPY --from=builder /app/build/eth_pokhar ./
COPY --from=builder /app/db/migrations ./db/migrations

ENTRYPOINT ["sh", "-c"]
