package selectpattern

import "time"

const (
	defaultTimeout = 10
	defaultCaching = false
)

type Connection struct {
	addr    string
	cache   bool
	timeout time.Duration
}

type options struct {
	timeout time.Duration
	caching bool
}

type Option interface {
	apply(*options)
}

type optionFunc func(*options)

func (f optionFunc) apply(o *options) {
	f.apply(o)
}
