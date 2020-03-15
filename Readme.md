
# DNS Server using blockchain

A PoC of a DNS server that uses blockchain (Ethereum) to store records.
Supports A, AAAA and CNAME records.

## Getting it going

### Prerequisites
- Go (duh!)
- Docker

### Installing Ethereum tools

Run the following to install Ethereum tools:
```
sudo apt-get install software-properties-common
sudo add-apt-repository -y ppa:ethereum/ethereum
sudo apt-get update
sudo apt-get install ethereum
sudo apt-get install solc
```

### Running the solution

First init the go modules by running this from the root folder of the repo:
```
go mod download
```

Then go to __config__ folder and create a __.env__ file within it with the following content:
```
CONTRACTADDR=""
GATEWAY="http://0.0.0.0:8545"
KEYSTOREFILE=""
KEYSTOREPASS=""
```

Stay in __config__ folder and run the following to create a new Ethereum wallet
```
geth --datadir . account new
``` 

Next you will have to get some ether, make note of your public Ethereum wallet address and go to [https://faucet.rinkeby.io/](https://faucet.rinkeby.io/) to get some.

Set __KEYSTOREFILE__ value in __.env__ to the name of the secret key file and __KEYSTOREPASS__ to the password you set for your Ethereum account, e.g.:
```
KEYSTOREFILE="UTC--2020-03-15T14-12-07.699399000Z--XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
KEYSTOREPASS="<password>"
```

Now you are ready to launch a light geth node client connected to rinkeby test network. Simply run this bash script from the root of the solution:
```
./launch-geth-node.sh
```

The node will take a few minutes to sync to latest blockchain state, once it does you are ready to test your setup.

Keep it running and from another terminal window go to __registrar-client__ folder and run the registrar client:
```
go run main.go
```

You will be greeted with a prompt like:
```
Pick an option:
1. Set record.
2. Show record.
3. Exit.
4. Reset and exit.
```

Pick _Set record_, type in _A_ when asked _Type in the record name_ type _A_, then enter the record name (__NOTE__: always end the domain name with a __.__, e.g.: __example.com.__) and finally an IP value for the record.

Wait until transaction is processed by the network and then choose _Show record_, type in the record type and name and you should see its value.

### Running the DNS Server

You could start the DNS server locally, by running it from __dns-server__ folder. But, to make it easier there is a docker-compose file included.

First, build the Docker image:
```
docker build -t blockchain-dns-server .
```

Then:
```
docker-compose up -d
docker exec -it blockchain-dns-server bash
```

Give it a few minutes for the geth node running inside the container to sync to latest network state then you can add records by running __registrar-client__. To test things you could use commands like __dig__, __wget__. Keep in mind that the DNS Server only supports A, AAAA and CNET records so its usage is a bit limited.

### A note on contract compilation

If you fancy changing the contract a bit here is a quick primer. The contract for this solution is in __blockchain-contract__ folder, once you modify the contract (_inet-dns-record.sol_) you have to compile it:
```
solc --abi --bin inet-dns-record.sol -o build
```

And then create a go proxy that you can use in your code:
```
abigen --abi="build/InetDnsRecord.abi" --bin="build/InetDnsRecord.bin" --pkg=dnsrecord --out="inet-dns-record.go"
```

