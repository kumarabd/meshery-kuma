package config

import (
	"github.com/layer5io/gokit/utils"
)

// Local instance for configuration
type Local struct {
	store map[string]string
}

// NewLocal intializes a local instance and dependencies
func NewLocal() (Handler, error) {
	return &Local{
		store: make(map[string]string),
	}, nil
}

// -------------------------------------------Application config methods----------------------------------------------------------------

// SetKey sets a key value in local store
func (l *Local) SetKey(key string, value string) error {
	l.store[key] = value
	return nil
}

// GetKey gets a key value from local store
func (l *Local) GetKey(key string) (string, error) {
	return l.store[key], nil
}

// Server provides server specific configuration
func (l *Local) Server(result interface{}) error {
	s, err := utils.Marshal(server)
	if err != nil {
		return ErrLocal(err)
	}
	return utils.Unmarshal(s, result)
}

// MeshSpec provides mesh specific configuration
func (l *Local) MeshSpec(result interface{}) error {
	s, err := utils.Marshal(mesh)
	if err != nil {
		return ErrLocal(err)
	}
	return utils.Unmarshal(s, result)
}

// MeshInstance provides mesh specific configuration
func (l *Local) MeshInstance(result interface{}) error {
	s, err := utils.Marshal(instance)
	if err != nil {
		return ErrLocal(err)
	}
	return utils.Unmarshal(s, result)
}

// Operations provides operations in the mesh
func (l *Local) Operations(result interface{}) error {
	d, err := utils.Marshal(operations)
	if err != nil {
		return err
	}
	return utils.Unmarshal(d, result)
}
