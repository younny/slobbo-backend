FROM golang:alpine

LABEL maintainer="Slobbo"

# Install git - required for fetching dependencies.
RUN apk update && apk add --no-cache git && apk add --no-cach bash && apk add build-base

# Set directory
RUN mkdir /app
WORKDIR /app

COPY . .
COPY .env .

# Download dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

# Build the Go app
RUN go build -o /build

EXPOSE 8080

CMD ["/build"]