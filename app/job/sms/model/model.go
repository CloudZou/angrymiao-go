package model

import (
	"container/ring"
	"encoding/json"
	"sync"
)

// Message .
type Message struct {
	Action string          `json:"action"`
	Table  string          `json:"table"`
	New    json.RawMessage `json:"new"`
	Old    json.RawMessage `json:"old"`
}

// ConcurrentRing thread-safe ring
type ConcurrentRing struct {
	*ring.Ring
	sync.Mutex
}
// NewConcurrentRing .
func NewConcurrentRing(length int) *ConcurrentRing {
	return &ConcurrentRing{Ring: ring.New(length)}
}
