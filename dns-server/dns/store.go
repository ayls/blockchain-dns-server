package dns

import (	
	"log"
	"context"
	"errors"
	"net"	

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
	envLoc = "../config/.env"	
	ErrTransactionWait = "If you've just started the application, wait a while for the network to confirm your transaction."
)

var (
	errTypeNotSupported = errors.New("Type not supported")
	errIPInvalid      = errors.New("Invalid IP address")
)

func (s *store) loadEnv() {
	var err error
	if myenv, err = godotenv.Read(envLoc); err != nil {
		s.logger.Printf("Could not load env from %s: %v", envLoc, err)
	}
}

func (s *store) updateEnvFile(k string, val string) {
	myenv[k] = val
	err := godotenv.Write(myenv, envLoc)
	if err != nil {
		s.logger.Printf("Failed to update %s: %v\n", envLoc, err)
	}
}

func (s *store) get(q dnsmessage.Question) ([]dnsmessage.Resource, bool) {
	recName := toStringQuestion(q)
	recType := uint16(q.Type)
	s.logger.Printf("Looking up %s, type %d", recName, recType)

	none := dnsmessage.Resource{}	
	recValue, err := s.session.GetRecord(recName, recType)
	if err != nil {
		s.logger.Printf("Can't answer lookup for %s, type %d: %v\n", recName, recType, err)		
		s.logger.Println(ErrTransactionWait)
		return []dnsmessage.Resource{none}, false
	}
	resource, err := toResource(recType, recName, recValue)
	if err != nil {
		s.logger.Printf("Can't answer lookup for %s, type %d: %v\n", recName, recType, err)				
		return []dnsmessage.Resource{none}, false
	}

	s.logger.Printf("Anwered lookup for %s, type %d", recName, recType)
	return []dnsmessage.Resource{resource}, true
}

func toStringQuestion(q dnsmessage.Question) string {
	b := make([]byte, q.Name.Length)
	for i := 0; i < int(q.Name.Length); i++ {
		b[i] = q.Name.Data[i]
	}

	return string(b)
}

func toResource(recType uint16, recName string, recValue string) (dnsmessage.Resource, error) {
	rName, err := dnsmessage.NewName(recName)
	none := dnsmessage.Resource{}
	if err != nil {
		return none, err
	}

	var rType dnsmessage.Type
	var rBody dnsmessage.ResourceBody

	switch dnsmessage.Type(recType) {
	case dnsmessage.TypeA:
		rType = dnsmessage.TypeA
		ip := net.ParseIP(recValue)
		if ip == nil {
			return none, errIPInvalid
		}
		rBody = &dnsmessage.AResource{A: [4]byte{ip[12], ip[13], ip[14], ip[15]}}
	case dnsmessage.TypeAAAA:
		rType = dnsmessage.TypeAAAA
		ip := net.ParseIP(recValue)
		if ip == nil {
			return none, errIPInvalid
		}
		var ipV6 [16]byte
		copy(ipV6[:], ip)
		rBody = &dnsmessage.AAAAResource{AAAA: ipV6}		
	case dnsmessage.TypeCNAME:
		rType = dnsmessage.TypeCNAME
		cname, err := dnsmessage.NewName(recValue)
		if err != nil {
			return none, err
		}
		rBody = &dnsmessage.CNAMEResource{CNAME: cname}
	case dnsmessage.TypeNS:
		fallthrough
	case dnsmessage.TypePTR:
		fallthrough
	case dnsmessage.TypeSOA:
		fallthrough	
	case dnsmessage.TypeMX:
		fallthrough
	case dnsmessage.TypeSRV:
		fallthrough		
	case dnsmessage.TypeTXT:
		fallthrough
	case dnsmessage.TypeOPT:
		fallthrough		
	default:
		return none, errTypeNotSupported
	}

	return dnsmessage.Resource{
		Header: dnsmessage.ResourceHeader{
			Name:  rName,
			Type:  rType,
			Class: dnsmessage.ClassINET,
			TTL:   0,
		},
		Body: rBody,
	}, nil
}

func (s *store) init() (*ethclient.Client) {
	s.loadEnv()

	// Load and init variables
	ctx := context.Background()

	// Connect to Ethereum gateway
	client, err := ethclient.Dial(myenv["GATEWAY"])
	s.client = client
	if err != nil {
		s.logger.Fatalf("Could not connect to Ethereum gateway: %v\n", err)
	}

	s.logger.Println("Connected to Ethereum gateway")

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
	session := contract.NewSession(ctx, "../config/keystore/" + myenv["KEYSTOREFILE"], myenv["KEYSTOREPASS"])
	s.session = session
	s.logger.Println("Contract session created")
}

func (s *store) newContract() {
	if myenv["CONTRACTADDR"] == "" {
		session, contractAddress := contract.NewContract(s.session, s.client)	
		s.session = session
		s.updateEnvFile("CONTRACTADDR", contractAddress)
		s.logger.Println("New contract created")
	}
}

func (s *store) loadContract() {
	session := contract.LoadContract(s.session, s.client, myenv["CONTRACTADDR"])
	s.session = session
	s.logger.Println("Contract loaded")	
}