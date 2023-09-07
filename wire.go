//go:build wireinject

package main

import "github.com/google/wire"

// initialize provides dependency injection process by the "google/wire" package.
func initialize() (*App, error) {
	wire.Build(newConfig, newApp)
	return &App{}, nil
}
