package main

import (
	"github.com/EliasStar/BacoTell/pkg/provider"
)

type TestComponent struct{}

var _ provider.Component = TestComponent{}

// CustomId implements provider.Component
func (TestComponent) CustomId() (string, error) {
	return "test_cpt", nil
}

// Handle implements provider.Component
func (TestComponent) Handle(provider.HandleProxy) error {
	logger.Info("handle component")
	return nil
}
