FROM golang:alpine as builder

LABEL maintainer="Slobbo <slobbodibo@gmail.com>"

# Install git - required for fetching dependencies.
RUN apk update && apk add --no-cache git && apk add --no-cach bash && apk add build-base

# Set directory
RUN mkdir /app
WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Start a new stage from scratch
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/
COPY --from=builder /app/main .
COPY --from=builder /app/.env .   


EXPOSE 8080

CMD ["./main"]