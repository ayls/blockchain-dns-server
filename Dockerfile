FROM golang:latest

RUN apt-get update
RUN apt-get install dnsutils -y

# build registrar-client
WORKDIR /app/registrar-client
COPY ./registrar-client ./
RUN go mod download
RUN go build -o main .

# build dns-server
WORKDIR /app/dns-server
COPY ./dns-server ./
RUN go mod download
RUN go build -o main .

EXPOSE 53

ENTRYPOINT ["./main"]