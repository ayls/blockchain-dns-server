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
	book   store
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
		s.logger.Println("Got query")		
		go s.Query(Packet{*addr, m})
	}
}

// Query lookup answers for DNS message.
func (s *DNSService) Query(p Packet) {

	// was checked before entering this routine
	q := p.message.Questions[0]

	// answer the question
	key := qString(q)
	s.logger.Printf("Looking up %s, %d", key, q.Type)
	val, ok := s.book.get(key)

	if ok {
		s.logger.Printf("Anwered lookup for %s", key)
		p.message.Answers = append(p.message.Answers, val...)
		go s.sendPacket(p.message, p.addr)
	} else {
		s.logger.Printf("Can't answer lookup for %s", key)
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
		book:   store{data: make(map[string]entry), logger: logger},
		logger: logger,
	}
}

// Start conveniently init every parts of DNS service.
func Start(logger *log.Logger) *DNSService {
	s := New(logger)
	s.book.load()
	s.Listen()

	return &s
}
