FROM golang:1.19.1-alpine3.16

WORKDIR /app

COPY *.mod *.sum ./
RUN go mod download

COPY ./ ./
RUN go build -o ./app ./cmd/main.go

EXPOSE 8080

RUN addgroup -S noroot && adduser -S noroot -G noroot
RUN chmod 766 /app/logs/*.txt
USER noroot:noroot

CMD ["./app"]