package main

import (
	"context"
	"net/http"
)

type Config struct {
	Methods []string `json:"Methods,omitempty" toml:"Methods,omitempty" yaml:"Methods,omitempty" export:"true"`
	Message string   `json:"Message,omitempty" toml:"Message,omitempty" yaml:"Message,omitempty" export:"true"`
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{
		Methods: []string{},
		Message: "",
	}
}

type MethodWhitelist struct {
	cfg  *Config
	next http.Handler
}

func New(_ context.Context, next http.Handler, config *Config, _ string) (http.Handler, error) {
	return &MethodWhitelist{
		cfg:  config,
		next: next,
	}, nil
}

func (m *MethodWhitelist) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	allowed := false
	for _, method := range m.cfg.Methods {
		if req.Method == method {
			allowed = true
			break
		}
	}
	if allowed {
		m.next.ServeHTTP(rw, req)
	} else {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		msg := m.cfg.Message
		if msg == "" {
			msg = "Method Not Allowed"
		}
		_, _ = rw.Write([]byte(msg))
	}
}
