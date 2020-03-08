package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"

	"ayls/blockchain-dns-server/registrar-client/dnsrecord"
)

var myenv map[string]string

const (
	envLoc = ".env" // Define location of env file to load here.
	// ErrTransactionWait should be returned/printed when we encounter an error that may be a result of the transaction not being confirmed yet.
	ErrTransactionWait = "if you've just started the application, wait a while for the network to confirm your transaction."
)

// loadEnv loads environment variables from location envLoc
// Call this at the top of every function that uses environment variables.
func loadEnv() {
	var err error
	if myenv, err = godotenv.Read(envLoc); err != nil {
		log.Printf("could not load env from %s: %v", envLoc, err)
	}
}

func main() {
	loadEnv()

	// Load and init variables
	ctx := context.Background()

	// Connect to Ethereum gateway
	client, err := ethclient.Dial(myenv["GATEWAY"])
	if err != nil {
		log.Fatalf("could not connect to Ethereum gateway: %v\n", err)
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

		// Reads a single UTF-8 character (rune)
		// from STDIN and switches to case.
		switch readStringStdin() {
		case "1":
			setRecord(session)
			break
		case "2":
			showRecord(session)
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

//// Contract initialization functions

// NewContract deploys a contract if no existing contract exists
func NewContract(session dnsrecord.DnsrecordSession, client *ethclient.Client) dnsrecord.DnsrecordSession {
	loadEnv()

	// Test our inputs
	if myenv["CONTRACTADDR"] != "" {
		return session
	}

	// Hash answer before sending it over Ethereum network.
	contractAddress, tx, instance, err := dnsrecord.DeployDnsrecord(&session.TransactOpts, client)
	if err != nil {
		log.Fatalf("could not deploy contract: %v\n", err)
	}
	fmt.Printf("Contract deployed! Wait for tx %s to be confirmed.\n", tx.Hash().Hex())

	session.Contract = instance
	updateEnvFile("CONTRACTADDR", contractAddress.Hex())
	return session
}

// LoadContract loads a contract if one exists
func LoadContract(session dnsrecord.DnsrecordSession, client *ethclient.Client) dnsrecord.DnsrecordSession {
	loadEnv()

	if myenv["CONTRACTADDR"] == "" {
		log.Println("could not find a contract address to load")
		return session
	}
	addr := common.HexToAddress(myenv["CONTRACTADDR"])
	instance, err := dnsrecord.NewDnsrecord(addr, client)
	if err != nil {
		log.Fatalf("could not load contract: %v\n", err)
		log.Println(ErrTransactionWait)
	}
	session.Contract = instance
	return session
}

// NewSession returns a quiz.QuizSession struct that
// contains an authentication key to sign transactions with.
func NewSession(ctx context.Context) (session dnsrecord.DnsrecordSession) {
	loadEnv()

	// Create new transactor
	keystore, err := os.Open(myenv["KEYSTORE"])
	if err != nil {
		log.Printf(
			"could not load keystore from location %s: %v\n",
			myenv["KEYSTORE"],
			err,
		)
	}
	defer keystore.Close()

	auth, err := bind.NewTransactor(keystore, myenv["KEYSTOREPASS"])
	if err != nil {
		log.Printf("%s\n", err)
	}

	// bind.NewTransactor() returns a bind.TransactOpts{} struct with the following field values:
	// From: auth.From,
	// Signer: auth.Signer,
	// Nonce: nil // Setting to nil uses nonce of pending state
	// Value: big.NewInt(0), // 0 because we're not transferring Eth
	// GasPrice: nil // Setting to nil automatically suggests a gas price
	// GasLimit: 0 // Setting to 0 automatically estimates gas limit

	// Return session without contract instance
	return dnsrecord.DnsrecordSession{
		TransactOpts: *auth,
		CallOpts: bind.CallOpts{
			From:    auth.From,
			Context: ctx,
		},
	}
}

//// Contract interaction functions

// setRecord sets a test record
func setRecord(session dnsrecord.DnsrecordSession) {
	// Send answer
	txSendAnswer, err := session.AddRecord("example.com", "93.184.216.34")
	if err != nil {
		log.Printf("could not set record in contract: %v\n", err)
		return
	}
	fmt.Printf("Record set! Please wait for tx %s to be confirmed.\n", txSendAnswer.Hash().Hex())
	return
}

// showRecord prints out the set record.
func showRecord(session dnsrecord.DnsrecordSession) {
	ip, err := session.GetRecord("example.com")
	if err != nil {
		log.Printf("could not read record from contract: %v\n", err)
		log.Println(ErrTransactionWait)
		return
	}
	fmt.Printf("Ip: %s\n", ip)
	return
}

//// Utility functions

// updateEnvFile saves the contract address to our .env file
func updateEnvFile(k string, val string) {
	myenv[k] = val
	err := godotenv.Write(myenv, envLoc)
	if err != nil {
		log.Printf("failed to update %s: %v\n", envLoc, err)
	}
}

// readStringStdin reads a string from STDIN and strips any trailing \n characters from it.
func readStringStdin() string {
	reader := bufio.NewReader(os.Stdin)
	inputVal, err := reader.ReadString('\n')
	if err != nil {
		log.Printf("invalid option: %v\n", err)
		return ""
	}

	output := strings.TrimSuffix(inputVal, "\n") // Important!
	return output
}
