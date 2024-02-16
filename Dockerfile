FROM golang:1.21-alpine as builder
RUN mkdir /app
WORKDIR /app
ADD . .

RUN go get
RUN go build -o ./build/eth_pokhar


FROM alpine:latest  
WORKDIR /
COPY --from=builder /app/build/eth_pokhar ./
COPY --from=builder /app/db/migrations ./db/migrations

ENTRYPOINT ["sh", "-c"]
