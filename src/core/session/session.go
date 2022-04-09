package session

import (
	"context"
	"net/http"
	"sync"
	"time"

	"github.com/astaxie/beego/cache/redis"
	"github.com/astaxie/beego/session"
	"github.com/cowk8s/harbor/src/lib/cache"
	"github.com/cowk8s/harbor/src/lib/log"
)

const (
	// HarborProviderName is the harbor session provider name
	HarborProviderName = "harbor"
)

var harborpder = &Provider{}

// SessionStore redis session store
type SessionStore struct {
	c           cache.Cache
	sid         string
	lock        sync.RWMutex
	values      map[interface{}]interface{}
	maxlifetime int64
}

// Set value in redis session
func (rs *SessionStore) Set(key, value interface{}) error {
	rs.lock.Lock()
	defer rs.lock.Unlock()
	rs.values[key] = value
	return nil
}

// Get value in redis session
func (rs *SessionStore) Get(key interface{}) interface{} {
	rs.lock.RLock()
	defer rs.lock.RUnlock()
	if v, ok := rs.values[key]; ok {
		return v
	}
	return nil
}

func (rs *SessionStore) Delete(key interface{}) error {
	rs.lock.Lock()
	defer rs.lock.Unlock()
	delete(rs.values, key)
	return nil
}

// Flush clear all values in redis session
func (rs *SessionStore) Flush() error {
	rs.lock.Lock()
	defer rs.lock.Unlock()
	rs.values = make(map[interface{}]interface{})
	return nil
}

// SessionID get redis session id
func (rs *SessionStore) SessionID() string {
	return rs.sid
}

func (rs *SessionStore) SessionRelease(w http.ResponseWriter) {
	b, err := session.EncodeGob(rs.values)
	if err != nil {
		return
	}

	if rdb, ok := rs.c.(*redis.Cache); ok {
		cmd := rdb.Client.Set(context.TODO(), rs.sid, string(b), time.Duration(rs.maxlifetime))
		if cmd.Err() != nil {
			log.Debugf("release session error: %v", err)
		}
	}
}

/// Provider redis session provider
type Provider struct {
	maxlifetime int64
	c           cache.Cache
}


