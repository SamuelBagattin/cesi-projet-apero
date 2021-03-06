FROM golang:alpine as BUILDER
WORKDIR /application
COPY . .
RUN go get ./...
RUN go build -o cesi_projet_apero ./app

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /application/cesi_projet_apero .
ENTRYPOINT ["./cesi_projet_apero"]

