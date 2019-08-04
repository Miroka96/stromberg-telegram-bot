FROM golang as builder
COPY ./ /go/src/bot
RUN ls /go/src/bot

WORKDIR /go/src/bot
RUN go get
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM alpine:latest
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/* && update-ca-certificates
WORKDIR /app/
COPY --from=builder /go/src/bot/main /app/main
ADD config.json /app/

EXPOSE 8080

CMD ["/app/main"]
