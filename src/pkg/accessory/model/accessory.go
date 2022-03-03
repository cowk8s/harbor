package model

const (
	RefNone = "base"
	RefSoft = "soft"
	RefHard = "hard"
)

type RefProvider interface {
	Kind() string
}
