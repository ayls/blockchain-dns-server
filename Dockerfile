# build
FROM golang:latest AS builder

WORKDIR /app
COPY . .
RUN go mod download

WORKDIR /app/registrar-client
RUN go build -o registrar-client .

WORKDIR /app/dns-server
RUN go build -o dns-server .

# final image
FROM ubuntu:latest

RUN apt-get update
RUN apt-get install dnsutils -y
RUN apt-get install wget -y 
RUN apt-get install software-properties-common -y 
RUN add-apt-repository ppa:ethereum/ethereum -y
RUN apt-get update
RUN apt-get install ethereum -y

WORKDIR /app

COPY rinkeby.json .
COPY docker-entrypoint.sh .
COPY --from=builder /app/registrar-client/registrar-client .
COPY --from=builder /app/dns-server/dns-server .

ENV PATH="/app:${PATH}"

EXPOSE 53

ENTRYPOINT ["./docker-entrypoint.sh"]