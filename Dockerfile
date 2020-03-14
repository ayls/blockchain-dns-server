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
FROM golang:latest

RUN apt-get update
RUN apt-get install dnsutils -y
RUN apt-get install -y wget

WORKDIR /app
ARG BINARY="geth-linux-amd64-1.9.11-6a62fe39.tar.gz"
RUN wget "https://gethstore.blob.core.windows.net/builds/$BINARY"
RUN tar -xzvf $BINARY --strip 1
RUN rm $BINARY

COPY *.sh ./
COPY rinkeby.json .
COPY --from=builder /app/registrar-client/registrar-client .
COPY --from=builder /app/dns-server/dns-server .

ENV PATH="/app:${PATH}"

EXPOSE 53

ENTRYPOINT ["./docker-entrypoint.sh"]