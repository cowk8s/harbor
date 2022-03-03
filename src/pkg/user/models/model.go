package models

// UserTable is the name of table in DB that holds the user object
const UserTable = "harbor_user"

// Option ...
type Option func(*Options)

// Options ...
type Options struct {
	IncludeDefaultAdmin bool
}

// WithDefaultAdmin set the IncludeAdmin = true
func WithDefaultAdmin() Option {
	return func(o *Options) {
		o.IncludeDefaultAdmin = true
	}
}

// NewOptions ...
func NewOptions(options ...Option) *Options {
	opts := &Options{}
	for _, f := range options {
		f(opts)
	}
	return opts
}
