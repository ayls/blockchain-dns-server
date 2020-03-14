package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"errors"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	
	"ayls/blockchain-dns-server/blockchain-contract"
	"ayls/blockchain-dns-server/blockchain-contract/dnsrecord"	
)

var myenv map[string]string

const (
	envLoc = "../config/.env"
	ErrTransactionWait = "If you've just started the application, wait a while for the network to confirm your transaction."
)

func loadEnv() {
	var err error
	if myenv, err = godotenv.Read(envLoc); err != nil {
		log.Printf("Could not load env from %s: %v", envLoc, err)
	}
}

func updateEnvFile(k string, val string) {
	myenv[k] = val
	err := godotenv.Write(myenv, envLoc)
	if err != nil {
		log.Printf("Failed to update %s: %v\n", envLoc, err)
	}
}

func main() {
	loadEnv()

	// Load and init variables
	ctx := context.Background()

	// Connect to Ethereum gateway
	client, err := ethclient.Dial(myenv["GATEWAY"])
	if err != nil {
		log.Fatalf("Could not connect to Ethereum gateway: %v\n", err)
	}
	defer client.Close()

	// Init new authenticated session
	session := NewSession(ctx)

	// Load or Deploy contract, and update session with contract instance
	if myenv["CONTRACTADDR"] == "" {
		session = NewContract(session, client)
	}

	// If we have an existing contract, load it; if we've deployed a new contract, attempt to load it.
	if myenv["CONTRACTADDR"] != "" {
		session = LoadContract(session, client)
	}

	// Loop to implement simple CLI
	for {
		fmt.Printf(
			"Pick an option:\n" + "" +
				"1. Set record.\n" +
				"2. Show record.\n" +
				"3. Exit.\n" +
				"4. Reset and exit.\n",
		)

		// Read selection
		switch readStringStdin() {
		case "1":
			fmt.Println("Type in the record type")
			recType, err := parseRecordType(readStringStdin())
			if err == nil {
				fmt.Println("Type in the record name")
				recName := readStringStdin()
				fmt.Println("Type in the record value")
				recValue := readStringStdin()
				setRecord(session, recType, recName, recValue)
			} else {
				log.Printf("%v\n", err)
			}
			break
		case "2":
			fmt.Println("Type in the record type")
			recType, err := parseRecordType(readStringStdin())
			if err == nil {
				fmt.Println("Type in the record name")
				recName := readStringStdin()
				showRecord(session, recType, recName)
			} else {
				log.Printf("%v\n", err)
			}			
			break
		case "3":
			fmt.Println("Bye!")
			return
		case "4":
			fmt.Println("Cleared contract address. Bye!")
			updateEnvFile("CONTRACTADDR", "")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
			break
		}
	}
}

func NewSession(ctx context.Context) (session dnsrecord.DnsrecordSession) {
	return contract.NewSession(ctx, "../config/keystore/" + myenv["KEYSTOREFILE"], myenv["KEYSTOREPASS"])
}

func NewContract(session dnsrecord.DnsrecordSession, client *ethclient.Client) dnsrecord.DnsrecordSession {
	if myenv["CONTRACTADDR"] != "" {
		return session
	}

	session, contractAddress := contract.NewContract(session, client)	
	updateEnvFile("CONTRACTADDR", contractAddress)
	return session
}

func LoadContract(session dnsrecord.DnsrecordSession, client *ethclient.Client) dnsrecord.DnsrecordSession {
	return contract.LoadContract(session, client, myenv["CONTRACTADDR"])
}

// readStringStdin reads a string from STDIN and strips any trailing \n characters from it.
func readStringStdin() string {
	reader := bufio.NewReader(os.Stdin)
	inputVal, err := reader.ReadString('\n')
	if err != nil {
		log.Printf("Invalid option: %v\n", err)
		return ""
	}

	output := strings.TrimSuffix(inputVal, "\n") // Important!
	return output
}

// parse record type input
func parseRecordType(recTypeString string) (uint16, error) {
	switch recTypeString {
	case "A":		
		return 1, nil
	case "NS":
		return 2, nil
	case "CNAME":
		return 5, nil
	case "SOA":
		return 6, nil
	case "PTR":
		return 12, nil
	case "MX":
		return 15, nil
	case "TXT":
		return 16, nil
	case "AAAA":
		return 28, nil
	case "SRV":
		return 33, nil
	default:
		return 0, errors.New("Unknown record type")		
	}
}

// setRecord sets a dns record
func setRecord(session dnsrecord.DnsrecordSession, recType uint16, recName string, recValue string) {
	// Send answer
	txSendAnswer, err := session.AddRecord(recName, recType, recValue)
	if err != nil {
		log.Printf("Could not set record in contract: %v\n", err)
		return
	}
	fmt.Printf("Record set! Please wait for tx %s to be confirmed.\n", txSendAnswer.Hash().Hex())
	return
}

// showRecord prints out a record.
func showRecord(session dnsrecord.DnsrecordSession, recType uint16, recName string) {
	ip, err := session.GetRecord(recName, recType)
	if err != nil {
		log.Printf("Could not read record from contract: %v\n", err)
		log.Println(ErrTransactionWait)
		return
	}
	fmt.Printf("Value: %s\n", ip)
	return
}