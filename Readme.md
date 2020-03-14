
docker build -t blockchain-dns-server .
docker rm -f dns-test
docker run -dit --mount source="$(PWD)/config",target=/config,type=bind --dns=0.0.0.0 --name dns-test blockchain-dns-server
docker exec -it dns-test bash


docker-compose up -d
docker exec -it blockchain-dns-server bash
docker-compose down

solc --abi --bin inet-dns-record.sol -o build
abigen --abi="build/InetDnsRecord.abi" --bin="build/InetDnsRecord.bin" --pkg=dnsrecord --out="inet-dns-record.go"



https://www.rinkeby.io/#geth

0xe5841D0117A81640385ea9A5550CD58f35a24794

geth --nousb --datadir=$HOME/.rinkeby init rinkeby.json


geth --nousb --networkid=4 --rpc --datadir=$HOME/.rinkeby --syncmode=light --ethstats='ayls:Respect my authoritah!@stats.rinkeby.io' --bootnodes=enode://a24ac7c5484ef4ed0c5eb2d36620ba4e4aa13b8c84684e1b4aab0cebea2ae45cb4d375b77eab56516d34bfbd3c1a833fc51296ff084b770b94fb9028c4d25ccf@52.169.42.101:30303

