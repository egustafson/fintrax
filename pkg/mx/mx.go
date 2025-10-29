// Package mx provides management extensions for monitoring components and their status.
package mx

import "strings"

// StatusObj represents the status of a component.
type StatusObj struct {
	Health  Status
	Details map[string]interface{}
}

// Statusable represents an object that can report its status.
type Statusable interface {
	TypeID() string
	Status() StatusObj
}

// MO represents a managed object
type MO interface {
	Statusable
	SetState(key string, value string)
	GetState(key string) (string, bool)
	Attach(name string, child Statusable)
}

// BaseMO is a basic implementation of the MO interface.
type BaseMO struct {
	State    map[string]string
	Children map[string]Statusable
}

// NewBaseMO creates a new BasicMO instance.
func NewBaseMO() MO {
	return &BaseMO{
		State:    make(map[string]string),
		Children: make(map[string]Statusable),
	}
}

// Status returns the status of the BasicMO.
func (b *BaseMO) Status() StatusObj {
	health := OK
	if h, ok := b.GetState("health"); ok {
		health = AsStatus(h)
	}
	for _, child := range b.Children {
		if child.Status().Health > health {
			health = child.Status().Health
		}
	}
	details := make(map[string]interface{})
	for k, v := range b.State {
		details[k] = v
	}
	return StatusObj{
		Health:  health,
		Details: details,
	}
}

// TypeID returns the type identifier of the BasicMO.
func (b *BaseMO) TypeID() string {
	if t, ok := b.GetState("type-id"); ok {
		return strings.TrimSpace(t)
	}
	return "basic-mo"
}

// SetState sets a detail value.
func (b *BaseMO) SetState(key string, value string) {
	b.State[key] = value
}

// GetState returns a detail.
func (b *BaseMO) GetState(key string) (string, bool) {
	val, ok := b.State[key]
	return val, ok
}

// Attach registers a child Statusable.
func (b *BaseMO) Attach(name string, child Statusable) {
	b.Children[name] = child
}
