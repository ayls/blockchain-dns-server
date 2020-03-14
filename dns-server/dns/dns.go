package dns

import (
	"golang.org/x/net/dns/dnsmessage"
	"log"
	"net"
)

// DNSServer will do Listen, Query and Send.
type DNSServer interface {
	Listen()
	Query(Packet)
}

// DNSService is an implementation of DNSServer interface.
type DNSService struct {
	conn   *net.UDPConn
	dnslookup   store
	logger *log.Logger
}

// Packet carries DNS packet payload and sender address.
type Packet struct {
	addr    net.UDPAddr
	message dnsmessage.Message
}

// Listen starts a DNS server on port 53
func (s *DNSService) Listen() {
	var err error
	conn, err := net.ListenUDP("udp", &net.UDPAddr{Port: 53})
	s.conn = conn

	if err != nil {
		s.logger.Fatal(err)
	}

	defer s.conn.Close()

	s.logger.Println("Listening on port 53")

	for {
		buf := make([]byte, 512)
		_, addr, err := s.conn.ReadFromUDP(buf)
		if err != nil {
			s.logger.Println(err)
			continue
		}
		var m dnsmessage.Message
		err = m.Unpack(buf)
		if err != nil {
			s.logger.Println(err)
			continue
		}
		if len(m.Questions) == 0 {
			continue
		}
		go s.query(Packet{*addr, m})
	}
}

// Query lookup answers for DNS message.
func (s *DNSService) query(p Packet) {

	// was checked before entering this routine
	q := p.message.Questions[0]

	// answer the question
	val, ok := s.dnslookup.get(q)

	if ok {
		p.message.Answers = append(p.message.Answers, val...)
		go s.sendPacket(p.message, p.addr)
	}
}

func (s *DNSService) sendPacket(message dnsmessage.Message, addr net.UDPAddr) {
	packed, err := message.Pack()
	if err != nil {
		s.logger.Println(err)
		return
	}

	_, err = s.conn.WriteToUDP(packed, &addr)
	if err != nil {
		s.logger.Println(err)
	}
}

// New setups a DNSService
func New(logger *log.Logger) DNSService {
	return DNSService{
		dnslookup:   store{logger: logger},
		logger: logger,
	}
}

// Start inits every part of DNS service.
func Start(logger *log.Logger) *DNSService {
	service := New(logger)
	client := service.dnslookup.init()
	defer client.Close()

	service.Listen()

	return &service
}
