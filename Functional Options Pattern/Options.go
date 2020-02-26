package main

import (
	"fmt"
	"time"
)

func main() {
	s := NewServices(
		SetName("peter"),
		SetTimeout(time.Second*5),
	)

	fmt.Println("name:", s.conf.Name)
	fmt.Println("time", s.conf.Timeout)
}

type Option func(options *Config)

type Config struct {
	Name    string
	Timeout time.Duration
}

type Services struct {
	conf Config
}

func SetTimeout(t time.Duration) Option {
	return func(options *Config) {
		options.Timeout = t
	}
}

func SetName(name string) Option {
	return func(options *Config) {
		options.Name = name
	}
}

func NewServices(opts ...Option) Services {
	c := Config{}
	for _, op := range opts {
		op(&c)
		fmt.Println("1")
	}
	s := Services{}
	s.conf = c
	return s
}