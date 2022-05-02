package uaa

import "sync"

// Auth is the implementation of AuthenicateHelper to access uaa for authentication.
type Auth struct {
	sync.Mutex
	client
}
