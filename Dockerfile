FROM golang:1.13-alpine as build

RUN apk add --no-cache git make gcc libc-dev

RUN go get github.com/julienschmidt/httprouter

RUN mkdir -p /app/factorial
COPY . /app/factorial
RUN cd /app/factorial && go build -o main .

FROM alpine:latest as runner

COPY --from=build /app/factorial/main /app/factorial/main

# Set the Current Working Directory inside the container
WORKDIR /app/factorial

# Expose port 8080 to the outside world
EXPOSE 8989

# Command to run the executable
CMD ["./main"]