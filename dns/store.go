package dns

import (
	"errors"
	"golang.org/x/net/dns/dnsmessage"
	"log"
	"sync"
	"time"
)

type store struct {
	sync.RWMutex
	logger *log.Logger
	data   map[string]entry
}

type entry struct {
	Resources []dnsmessage.Resource
	TTL       uint32
	Created   int64
}

type request struct {
	Host string
	TTL  uint32
	Type string
	Data string
}

var (
	errTypeNotSupport = errors.New("type not support")
	errIPInvalid      = errors.New("invalid IP address")
)

func (s *store) get(key string) ([]dnsmessage.Resource, bool) {
	s.RLock()
	e, ok := s.data[key]
	s.RUnlock()
	return e.Resources, ok
}

func (s *store) load() {
	s.set(request{
		Host: "example.com.",
		TTL:  600,
		Type: "A",
		Data: "93.184.216.34",
	})
	s.set(request{
		Host: "example.com.",
		TTL:  600,
		Type: "AAAA",
		Data: "2606:2800:220:1:248:1893:25c8:1946",
	})
}

func (s *store) set(req request) bool {
	changed := false
	resource, _ := toResource(req)
	key := ntString(resource.Header.Name, resource.Header.Type)
	s.logger.Printf("Storing data for %s", key)
	s.Lock()
	e := entry{
		Resources: []dnsmessage.Resource{resource},
		TTL:       resource.Header.TTL,
		Created:   time.Now().Unix(),
	}
	s.data[key] = e
	s.Unlock()

	return changed
}
