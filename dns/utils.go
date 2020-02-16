package dns

import (
	"golang.org/x/net/dns/dnsmessage"
	"net"
)

// question to string
func qString(q dnsmessage.Question) string {
	b := make([]byte, q.Name.Length+2)
	for i := 0; i < int(q.Name.Length); i++ {
		b[i] = q.Name.Data[i]
	}
	b[q.Name.Length] = uint8(q.Type >> 8)
	b[q.Name.Length+1] = uint8(q.Type)

	return string(b)
}

// resource name and type to string
func ntString(rName dnsmessage.Name, rType dnsmessage.Type) string {
	b := make([]byte, rName.Length+2)
	for i := 0; i < int(rName.Length); i++ {
		b[i] = rName.Data[i]
	}
	b[rName.Length] = uint8(rType >> 8)
	b[rName.Length+1] = uint8(rType)

	return string(b)
}

func toResource(req request) (dnsmessage.Resource, error) {
	rName, err := dnsmessage.NewName(req.Host)
	none := dnsmessage.Resource{}
	if err != nil {
		return none, err
	}

	var rType dnsmessage.Type
	var rBody dnsmessage.ResourceBody

	switch req.Type {
	case "A":
		rType = dnsmessage.TypeA
		ip := net.ParseIP(req.Data)
		if ip == nil {
			return none, errIPInvalid
		}
		rBody = &dnsmessage.AResource{A: [4]byte{ip[12], ip[13], ip[14], ip[15]}}
	case "AAAA":
		rType = dnsmessage.TypeAAAA
		ip := net.ParseIP(req.Data)
		if ip == nil {
			return none, errIPInvalid
		}
		var ipV6 [16]byte
		copy(ipV6[:], ip)
		rBody = &dnsmessage.AAAAResource{AAAA: ipV6}
	default:
		return none, errTypeNotSupport
	}

	return dnsmessage.Resource{
		Header: dnsmessage.ResourceHeader{
			Name:  rName,
			Type:  rType,
			Class: dnsmessage.ClassINET,
			TTL:   req.TTL,
		},
		Body: rBody,
	}, nil
}
