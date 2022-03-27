package cache

import (
	"time"
)

// Option function to set the options of the cache
type Option func(*Options)

// Options options of the cache
type Options struct {
	Address    string        // the address of the cache
	Codec      Codec         // the codec for the cache
	Expiration time.Duration // the default expiration for the cache
	Prefix     string        // the prefix for all the keys in the cache
}

// Key returns the real cache key
func (opts *Options) Key(key string) string {
	return opts.Prefix + key
}

func newOptions(opt ...Option) Options {
	opts := Options{}

	for _, o := range opt {
		o(&opts)
	}

	return opts
}

// Address sets the address
func Address(addr string) Option {
	return func(o *Options) {
		o.Address = addr
	}
}

// Expiration sets the default expiration
func Expiration(d time.Duration) Option {
	return func(o *Options) {
		o.Expiration = d
	}
}

// Prefix sets the prefix
func Prefix(prefix string) Option {
	return func(o *Options) {
		o.Prefix = prefix
	}
}
