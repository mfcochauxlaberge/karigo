package karigo

import (
	"sync"
	"time"
)

// Limiter ...
type Limiter struct {
	readTotals  map[string]uint
	writeTotals map[string]uint
	tokens      map[string]token

	sync.Mutex
}

type token struct {
	id         string
	user       string
	expiration time.Time
}
