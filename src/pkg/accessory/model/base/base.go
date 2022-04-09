package base

import (
	"github.com/cowk8s/harbor/src/pkg/accessory/model"
)

var _ model.Accessory = (*Default)(nil)

// Default default model with TypeNone and RefNone
type Default struct {
	Data model.AccessoryData
}

// Kind ...
func (a *Default) Kind() string {
	return model.RefNone
}

// IsSoft ...
func (a *Default) IsSoft() bool {
	return false
}

// IsHard ...
func (a *Default) IsHard() bool {
	return false
}

// Display ...
func (a *Default) Display() bool {
	return false
}

// GetData ...
func (a *Default) GetData() model.AccessoryData {
	return a.Data
}

// New returns base
func New(data model.AccessoryData) model.Accessory {
	return &Default{Data: data}
}

func init() {
	model.Register(model.TypeNone, New)
}
