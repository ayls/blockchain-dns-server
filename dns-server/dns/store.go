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
	errTypeNotSupport = errors.New("type not supported")
	errIPInvalid      = errors.New("invalid IP address")
)

func (s *store) get(key string) ([]dnsmessage.Resource, bool) {
	s.RLock()
	e, ok := s.data["ayls.dev"]
	s.RUnlock()
	return e.Resources, ok
}

func (s *store) load() {
	s.set(request{
		Host: "ayls.dev.",
		TTL:  600,
		Type: "A",
		Data: "51.144.90.155",
	})
}

func (s *store) set(req request) bool {
	changed := false
	resource, _ := toResource(req)
	key := "ayls.dev"
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
