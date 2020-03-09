FROM golang:latest

RUN apt-get update
RUN apt-get install dnsutils -y

WORKDIR /app
COPY . .
RUN go mod download

# build registrar-client
WORKDIR /app/registrar-client
RUN go build -o main .

# build dns-server
WORKDIR /app/dns-server
RUN go build -o main .

EXPOSE 53

# ENTRYPOINT ["./main"]