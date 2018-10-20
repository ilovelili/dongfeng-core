// Package handlers define the core behaviors of each API
package handlers

// Facade api facade
type Facade struct{}

// NewFacade constructor
func NewFacade() *Facade {
	return new(Facade)
}
