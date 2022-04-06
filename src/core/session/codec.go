package session

import (
	"encoding/gob"

	"github.com/astaxie/beego/session"
	commonmodels "github.com/cowk8s/harbor/src/common/models"
	"github.com/cowk8s/harbor/src/lib/cache"
	"github.com/cowk8s/harbor/src/lib/errors"
)

func init() {
	gob.Register(commonmodels.User{})
}

var (
	// codec the default codec for the cache
	codec cache.Codec = &gobCodec{}
)

type gobCodec struct{}

func (*gobCodec) Encode(v interface{}) ([]byte, error) {
	if vm, ok := v.(map[interface{}]interface{}); ok {
		return session.EncodeGob(vm)
	}

	return nil, errors.Errorf("object type invalid, %#v", v)
}

func (*gobCodec) Decode(data []byte, v interface{}) error {
	vm, err := session.DecodeGob(data)
	if err != nil {
		return err
	}

	switch in := v.(type) {
	case map[interface{}]interface{}:
		for k, v := range vm {
			in[k] = v
		}
	case *map[interface{}]interface{}:
		m := *in
		for k, v := range vm {
			m[k] = v
		}
	default:
		return errors.Errorf("object type invalid, %#v", v)
	}

	return nil
}
