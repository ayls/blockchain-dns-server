package dns

import (
	"net"
	"errors"

	"golang.org/x/net/dns/dnsmessage"
)

// question to string
func qString(q dnsmessage.Question) string {
	b := make([]byte, q.Name.Length)
	for i := 0; i < int(q.Name.Length); i++ {
		b[i] = q.Name.Data[i]
	}

	return string(b)
}

var (
	errTypeNotSupported = errors.New("type not supported")
	errIPInvalid      = errors.New("invalid IP address")
)

func toResource(recType int8, recName string, recValue string) (dnsmessage.Resource, error) {
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
