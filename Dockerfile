FROM golang as builder
RUN git clone https://github.com/Miroka96/stromberg-telegram-bot.git /go/src/bot

WORKDIR /go/src/bot
RUN go get
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM alpine:latest
WORKDIR /app/
COPY --from=builder /go/src/bot/main /app/main
ADD config.json /app/

EXPOSE 8080

CMD ["/bot/main"]
