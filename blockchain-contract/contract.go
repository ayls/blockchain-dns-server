package contract

import (
	"os"
	"log"
	"fmt"
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"	
	"github.com/ethereum/go-ethereum/common"	
	"github.com/ethereum/go-ethereum/ethclient"

	"ayls/blockchain-dns-server/blockchain-contract/dnsrecord"
)

const (
	// ErrTransactionWait should be returned/printed when we encounter an error that may be a result of the transaction not being confirmed yet.
	ErrTransactionWait = "if you've just started the application, wait a while for the network to confirm your transaction."
)

// NewContract deploys a contract
func NewContract(session dnsrecord.DnsrecordSession, client *ethclient.Client) (dnsrecord.DnsrecordSession, string) {
	// Hash answer before sending it over Ethereum network.
	contractAddress, tx, instance, err := dnsrecord.DeployDnsrecord(&session.TransactOpts, client)
	if err != nil {
		log.Fatalf("could not deploy contract: %v\n", err)
	}
	fmt.Printf("Contract deployed! Wait for tx %s to be confirmed.\n", tx.Hash().Hex())

	session.Contract = instance
	return session, contractAddress.Hex()
}

// LoadContract loads a contract if one exists
func LoadContract(session dnsrecord.DnsrecordSession, client *ethclient.Client, address string) dnsrecord.DnsrecordSession {
	if address == "" {
		log.Println("could not find a contract address to load")
		return session
	}
	addr := common.HexToAddress(address)
	instance, err := dnsrecord.NewDnsrecord(addr, client)
	if err != nil {
		log.Fatalf("could not load contract: %v\n", err)
		log.Println(ErrTransactionWait)
	}
	session.Contract = instance
	return session
}

// NewSession returns a quiz.DnsrecordSession struct that
// contains an authentication key to sign transactions with.
func NewSession(ctx context.Context, keystorePath string, keystorePass string) (session dnsrecord.DnsrecordSession) {
	// Create new transactor
	keystore, err := os.Open(keystorePath)
	if err != nil {
		log.Fatalf(
			"could not load keystore from location %s: %v\n",
			keystorePath,
			err,
		)
	}
	defer keystore.Close()

	auth, err := bind.NewTransactor(keystore, keystorePass)
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