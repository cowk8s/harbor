package inmemory

import (
	"context"
	"sync"
)

func init() {

}

type Driver struct {
	sync.Mutex
	cfgMap map[string]interface{}
}

func (d *Driver) Load(context.Context)
