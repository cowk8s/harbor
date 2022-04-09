package cosign

import (
	"github.com/cowk8s/harbor/src/pkg/accessory/model"
	"github.com/cowk8s/harbor/src/pkg/accessory/model/base"
)

// Signature signature model
type Signature struct {
	base.Default
}

// Kind gives the reference type of cosign signature.
func (c *Signature) Kind() string {
	return model.RefHard
}

// IsHard ...
func (c *Signature) IsHard() bool {
	return true
}

// New returns cosign signature
func New(data model.AccessoryData) model.Accessory {
	return &Signature{base.Default{
		Data: data,
	}}
}

func init() {
	model.Register(model.TypeCosignSignature, New)
}
