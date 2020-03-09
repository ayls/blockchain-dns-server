package dns

import (	
	"log"
	"context"

	"golang.org/x/net/dns/dnsmessage"	
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	
	"ayls/blockchain-dns-server/blockchain-contract"
	"ayls/blockchain-dns-server/blockchain-contract/dnsrecord"		
)

type store struct {
	logger *log.Logger
	client *ethclient.Client
	session dnsrecord.DnsrecordSession
}

var myenv map[string]string

const (
	envLoc = ".env" // Define location of env file to load here.
	// ErrTransactionWait should be returned/printed when we encounter an error that may be a result of the transaction not being confirmed yet.
	ErrTransactionWait = "if you've just started the application, wait a while for the network to confirm your transaction."
)

// loadEnv loads environment variables from location envLoc
// Call this at the top of every function that uses environment variables.
func (s *store) loadEnv() {
	var err error
	if myenv, err = godotenv.Read(envLoc); err != nil {
		s.logger.Printf("could not load env from %s: %v", envLoc, err)
	}
}

func (s *store) updateEnvFile(k string, val string) {
	myenv[k] = val
	err := godotenv.Write(myenv, envLoc)
	if err != nil {
		s.logger.Printf("failed to update %s: %v\n", envLoc, err)
	}
}

func (s *store) get(recType int8, recName string) ([]dnsmessage.Resource, bool) {
	none := dnsmessage.Resource{}	
	recValue, err := s.session.GetRecord(recName, recType)
	if err != nil {
		s.logger.Printf("could not read record from contract: %v\n", err)
		s.logger.Println(ErrTransactionWait)
		return []dnsmessage.Resource{none}, false
	}
	resource, err := toResource(recType, recName, recValue)
	if err != nil {
		return []dnsmessage.Resource{none}, false
	}
	return []dnsmessage.Resource{resource}, true
}

func (s *store) init() (*ethclient.Client) {
	s.loadEnv()

	// Load and init variables
	ctx := context.Background()

	// Connect to Ethereum gateway
	client, err := ethclient.Dial(myenv["GATEWAY"])
	s.client = client
	if err != nil {
		s.logger.Fatalf("could not connect to Ethereum gateway: %v\n", err)
	}

	// Init new authenticated session
	s.newSession(ctx)

	// Load or Deploy contract, and update session with contract instance
	if myenv["CONTRACTADDR"] == "" {
		s.newContract()
	}

	// If we have an existing contract, load it; if we've deployed a new contract, attempt to load it.
	if myenv["CONTRACTADDR"] != "" {
		s.loadContract()
	}	

	return s.client
}

func (s *store) newSession(ctx context.Context) {
	session := contract.NewSession(ctx, myenv["KEYSTORE"], myenv["KEYSTOREPASS"])
	s.session = session
}

func (s *store) newContract() {
	if myenv["CONTRACTADDR"] == "" {
		session, contractAddress := contract.NewContract(s.session, s.client)	
		s.session = session
		s.updateEnvFile("CONTRACTADDR", contractAddress)
	}
}

func (s *store) loadContract() {
	session := contract.LoadContract(s.session, s.client, myenv["CONTRACTADDR"])
	s.session = session
}